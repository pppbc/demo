<template>
	<div>
		<van-nav-bar title="书籍借还中心" fixed left-text="返回" left-arrow @click-left="onClickLeft" />
		<div style="margin-top:45px"></div>

		<grid :gridItems="gridItems" :columnNum="columnNum" :showClass="showClass" :gutter="gutter"></grid>

		<!-- 弹出层 -->

		<!-- 借书 -->
		<van-popup v-model="$store.state.lendbook" position="bottom" :style="{height:'80%'}" @click-overlay="clickoverlay">
			<div style="margin:20px 10px;color:gray;">借阅书籍</div>

			<van-search v-model="bookname" placeholder="请输入书名" show-action @search="onSearchBook">
				<van-icon slot="action" @click="onSearchBook">搜索图书</van-icon>
			</van-search>
			<van-search v-model="username" placeholder="请输入用户名" show-action @search="onSearchUser">
				<van-icon slot="action" @click="onSearchUser">搜索用户</van-icon>
			</van-search>
			<van-cell-group>
				<tbody id="bookinfo" class="body"></tbody>
				<tbody id="userinfo" class="body"></tbody>
			</van-cell-group>
			<input class="submit" @click="submitlend()" type="submit" value="确认借阅">
		</van-popup>
		<!-- 还书 -->
		<van-popup v-model="$store.state.returnbook" position="bottom" :style="{height:'80%'}" @click-overlay="clickoverlay">
			<div style="margin:20px 10px;color:gray;">归还书籍</div>
			<van-search v-model="username" placeholder="请输入用户名" show-action @search="onSearchListByUser">
				<van-icon slot="action" @click="onSearchListByUser">搜索用户借书信息</van-icon>
			</van-search>
			<van-search v-model="bookname" placeholder="请输入书名" show-action>
			</van-search>
			<table id="showbookinfo">
			</table>

			<input class="submit" @click="submitreturn()" type="submit" value="确认还书">
		</van-popup>
		<!-- 显示 -->
		<van-popup v-model="$store.state.showlist" position="bottom" :style="{height:'80%'}" @click-overlay="clickoverlay">
			<div style="margin:20px 10px;color:gray;">借书列表</div>
			<table>
				<thead>
					<tr class="top">
						<th width="60px">书名</th>
						<th width="70px">借书人</th>
						<th width="100px">借书时间</th>
						<th width="50px">状态</th>
						<th width="100px">还书时间</th>
					</tr>
				</thead>
				<tbody id="allinfo" class="body1"></tbody>
			</table>
		</van-popup>
	</div>
</template>

<script>
	import $ from 'jquery'
	import grid from '../components/grid.vue'
	export default {
		components: {
			grid,
		},
		data() {
			return {
				bookname: '',
				username: '',
				gridItems: [{
						text: '借书',
						path: '',
						icon: 'info-o',
						method: 'lendbook'
					},
					{
						text: '还书',
						path: '',
						icon: 'close',
						method: 'returnbook'
					},
					{
						text: '查看',
						path: '',
						icon: 'search',
						method: 'showlist'
					}
				],
				columnNum: 4,
				showClass: 'gridManage',
				gutter: 0,
				token:''
			}
		},
		beforeMount: function() {
			var info = localStorage.getItem("info")
			info = JSON.parse(info)
			this.token=info.token
		},
		methods: {
			onClickLeft() {
				this.$router.go(-1)
			},
			clickoverlay() {
				this.bookname = ''
				this.username = ''
				this.$store.state.returnbook = false
				this.$store.state.lendbook = false
				this.$store.state.showlist = false
				var html=''
				$("#showbookinfo").html(html)
				$("#bookinfo").html(html)
				$("#userinfo").html(html)
			},
			onSearchBook() {
				$.ajax({
					url: "http://47.98.227.139:9001/book/get",
					type: "get",
					data: {
						title: this.bookname,
						token:this.token
					},
					dataType: "json",
					success: (data) => {
						if (data.desc && data.desc === 'token异常，重新登陆') {
							this.$toast(data.desc);
							this.$router.push('/');
						}
						var json = data.data
						if (json) {
							var html = ''
							html += "<tr>";
							html += "<td width='90px'>" + "《" + json.title + "》" + "</td>";
							html += "<td width='90px'>" + json.author + "</td>";
							html += "<td width='50px'>" + json.price + "</td>";
							html += "<td width='50px'>" + json.num + "</td>";
							html += "<td width='100px'>" + json.desc + "</td>";
							html += "</tr>";
							$("#bookinfo").html(html)
							// window.console.log(json)
						}
						this.$toast(data.desc)
					},
					error: function() {
						alert("can't find server")
					}
				})
			},
			submitlend() {
				$.ajax({
					url: "http://47.98.227.139:9001/lend/lend",
					type: "post",
					data: {
						bookname: this.bookname,
						username: this.username,
						token:this.token
					},
					dataType: "json",
					success: (data) => {
						if (data.desc && data.desc === 'token异常，重新登陆') {
							this.$toast(data.desc);
							this.$router.push('/');
						}
						this.$toast(data.desc)
					},
					error: () => {
						alert("can't find server")
					}
				})
			},
			onSearchUser() {
				$.ajax({
					url: "http://47.98.227.139:9001/user/getname",
					type: "get",
					data: {
						username: this.username,
						token:this.token
					},
					dataType: "json",
					success: (data) => {
						if (data.desc && data.desc === 'token异常，重新登陆') {
							this.$toast(data.desc);
							this.$router.push('/');
						}
						var json = data.data
						if (json) {
							var html = ''
							html += "<tr>";
							html += "<td width='90px'>" + json.id + "</td>";
							html += "<td width='90px'>" + json.username + "</td>";
							html += "<td width='50px'>" + json.is_m + "</td>";
							html += "<td width='50px'>" + json.age + "</td>";
							html += "<td width='100px'>" + json.tel + "</td>";
							html += "</tr>";
							$("#userinfo").html(html)
						}
						this.$toast(data.desc)
					},
					error: function() {
						alert("can't find server")
					}
				})
			},
			submitreturn() {
				$.ajax({
					url: "http://47.98.227.139:9001/lend/return",
					type: "post",
					data: {
						bookname: this.bookname,
						username: this.username,
						token:this.token
					},
					dataType: "json",
					success: (data) => {
						if (data.desc && data.desc === 'token异常，重新登陆') {
							this.$toast(data.desc);
							this.$router.push('/');
						}
						this.$toast(data.desc)
					},
					error: () => {
						alert("can't find server")
					}
				})
			},
			onSearchListByUser() {
				$.ajax({
					url: "http://47.98.227.139:9001/lend/getbyuser",
					type: "get",
					data: {
						username: this.username,
						token:this.token
					},
					dataType: "json",
					success: (data) => {
						if (data.desc && data.desc === 'token异常，重新登陆') {
							this.$toast(data.desc);
							this.$router.push('/');
						}
						var json = data.data
						if(json){
							var html = ''
							
							html += "<thead>"
							html += "<tr>"
							html += "<th width='60px'>书名</th>"
							html += "<th width='70px'>借书人</th>"
							html += "<th width='100px'>借书时间</th>"
							html += "<th width='50px'>状态</th>"
							html += "<th width='100px'>还书时间</th>"
							html += "</tr>"
							html += "</thead>"
							html += '<tbody">';
							$.each(json, function(index, item) {
								html += "<tr class='top'>";
								html += "<td>" + item.username + "</td>";
								html += "<td>" + item.bookname + "</td>";
								html += "<td>" + item.lend_time + "</td>";
								html += "<td>" + item.is_r + "</td>";
								html += "<td>" + item.return_time + "</td>";
								html += "</tr>";
							})
							html += "</tbody>"
							$("#showbookinfo").html(html)
							html=''
							
						}else{
							html=''
							$("#showbookinfo").html(html)
						}
					},
					error: function() {
						alert("can't find server")
					}
				})
			}
		}
	}
</script>

<style scoped="scoped">
	.body {
		font-size: 14px;
		font-style: italic;
		line-height: 50px;
		color: #F08080;
	}

	.body1 {
		font-size: 14px;
		font-style: italic;
		line-height: 30px;
		color: #F08080;
	}
	#showbookinfo>tbody>tr>td{
		font-size: 14px;
		font-style: italic;
		line-height: 50px;
		color: red;
	}
	.submit{
		margin-top: 50px;
		border-radius: 10px;
		font-size: 16px;
		background-color: whitesmoke;
	}
</style>
