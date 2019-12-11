<template>
	<div>
		<!-- 顶部信息 -->
		<van-nav-bar title="书籍管理中心" fixed left-text="返回" left-arrow @click-left="onClickLeft" />

		<div style="margin-top:45px"></div>

		<!-- 功能版块 -->
		<grid :gridItems="gridItems" :columnNum="columnNum" :showClass="showClass" :gutter="gutter"></grid>

		<!-- 弹出层 -->

		<!-- 修改信息 -->
		<van-popup v-model="$store.state.showChangeB" position="bottom" :style="{height:'80%'}" @click-overlay="clickoverlay">
			<div style="margin:20px 10px;color:gray;">修改图书信息</div>
			<van-search v-model="bookname" placeholder="搜索书名" show-action @search="onSearch">
				<!-- slot="action"自定义插槽 -->
				<van-icon slot="action" @click="onSearch">搜索</van-icon>
			</van-search>
			<van-cell-group>
				<van-field v-model="title" required label="书名" right-icon="question-o" placeholder="请输入书名" />
				<van-field v-model="author" label="作者" placeholder="请输入作者" />
				<van-field v-model="num" type="number" label="剩余数量" placeholder="请输入余量" />
				<van-field v-model="price" type="number" label="价格" placeholder="请输入价格" />
				<van-field v-model="desc" label="书籍详情" type="textarea" placeholder="请输入书籍描述" rows="2" autosize />
			</van-cell-group>
			<input class="submit" @click="submitchange()" type="submit" value="确认修改">
		</van-popup>
		
		
		<!-- 添加书籍 -->
		
		
		<van-popup v-model="$store.state.showAddB" position="bottom" :style="{height:'70%'}" @click-overlay="clickoverlay">
			<div style="margin:20px 10px;color:gray;">添加书籍信息</div>
			<van-cell-group>
				<van-field v-model="title" required label="书名" right-icon="question-o" placeholder="请输入书名" />
				<van-field v-model="author" label="作者" placeholder="请输入作者" />
				<van-field v-model="num" type="number" label="剩余数量" placeholder="请输入余量" />
				<van-field v-model="price" type="number" label="价格" placeholder="请输入价格" />
				<van-field v-model="desc" label="书籍详情" type="textarea" placeholder="请输入书籍描述" rows="2" autosize />
			</van-cell-group>
			<input class="submit" @click="submitadd()" type="submit" value="确认添加">
		</van-popup>
		<!-- 删除书籍 -->
		<van-popup v-model="$store.state.showDeleteB" position="bottom" :style="{height:'80%'}" @click-overlay="clickoverlay">
			<div style="margin:20px 10px;color:gray;">删除图书信息</div>
			<van-search v-model="bookname" placeholder="搜索书名" show-action @search="onSearch">
				<!-- slot="action"自定义插槽 -->
				<van-icon slot="action" @click="onSearch">搜索</van-icon>
			</van-search>
			<van-cell-group>
				<van-field v-model="title" label="书名" placeholder="书名" disabled />
				<van-field v-model="author" label="作者" placeholder="作者" disabled />
				<van-field v-model="bookid" label="编号" placeholder="编号" disabled />
				<van-field v-model="num" type="number" label="剩余数量" placeholder="余量" disabled />
				<van-field v-model="price" type="number" label="价格" placeholder="价格" disabled />
				<van-field v-model="desc" label="书籍详情" type="textarea" placeholder="请输入书籍描述" rows="2" autosize disabled />
			</van-cell-group>
			<input class="submit" @click="submitdelete()" type="submit" value="确认删除">
		</van-popup>
		<!-- 搜索书籍 -->
		<van-popup v-model="$store.state.showSearchB" position="bottom" :style="{height:'70%'}" @click-overlay="clickoverlay">
			<div style="margin:20px 10px;color:gray;">搜索图书信息</div>
			<van-search v-model="bookname" placeholder="搜索书名" show-action @search="onSearch">
				<!-- slot="action"自定义插槽 -->
				<van-icon slot="action" @click="onSearch">搜索</van-icon>
			</van-search>
			<van-cell-group>
				<van-field v-model="title" label="书名" placeholder="书名" disabled />
				<van-field v-model="author" label="作者" placeholder="作者" disabled />
				<van-field v-model="bookid" label="编号" placeholder="编号" disabled />
				<van-field v-model="num" type="number" label="剩余数量" placeholder="余量" disabled />
				<van-field v-model="price" type="number" label="价格" placeholder="价格" disabled />
				<van-field v-model="desc" label="书籍详情" type="textarea" placeholder="请输入书籍描述" rows="2" autosize disabled />
			</van-cell-group>
		</van-popup>
	</div>
</template>

<script>
	import {
		Popup,
		Search,
		Icon,
		Toast
	} from 'vant'
	import $ from 'jquery'
	import grid from '@/components/grid.vue'
	export default {
		components: {
			grid,
			[Popup.name]: Popup,
			[Search.name]: Search,
			[Icon.name]: Icon,
			[Toast.name]: Toast
		},
		data() {
			return {
				allItems: Array,
				gridItems: [{
						text: '添加书籍',
						path: '',
						icon: 'add-o',
						method: 'add'
					},
					{
						text: '修改书籍',
						path: '',
						icon: 'info-o',
						method: 'change'
					},
					{
						text: '删除书籍',
						path: '',
						icon: 'close',
						method: 'delete'
					},
					{
						text: '搜索书籍',
						path: '',
						icon: 'search',
						method: 'search'
					},
					{
						text: '展示所有信息',
						path: '/justlook',
						icon: 'description'
					}
				],
				columnNum: 4,
				showClass: 'gridManage',
				gutter: 0,
				// 接收服务器传来的值
				title: '',
				author: '',
				num: 0,
				price: 0.0,
				bookname: '',
				desc: '',
				bookid: 0,
				booksrc: '',
				token:''
			}
		},
		beforeMount:function(){
			var info = localStorage.getItem("info")
			info = JSON.parse(info)
			this.token=info.token
		},
		methods: {
			clickoverlay() {
				this.$store.state.showChangeB = false
				this.$store.state.showAddB = false
				this.$store.state.showDeleteB = false
				this.$store.state.showSearchB = false
				this.title = ''
				this.author = ''
				this.price = ''
				this.num = ''
				this.desc = ''
				this.bookid = ''
				this.booksrc = ''
			},
			onClickLeft() {
				this.$router.go(-1)
			},
			onSearch() {
				//获取信息
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
							this.title = json.title
							this.author = json.author
							this.price = json.price
							this.num = json.num
							this.desc = json.desc
							this.bookid = json.id
							this.booksrc = json.src
						}
						window.console.log(json)
						this.$toast(data.desc)
					},
					error: function() {
						alert("can't find server")
					}
				})
			},
			submitchange() {
				$.ajax({
					url: "http://47.98.227.139:9001/book/change",
					type: "post",
					data: {
						id: this.bookid,
						title: this.title,
						author: this.author,
						price: this.price,
						num: this.num,
						desc: this.desc,
						src: this.booksrc,
						token:this.token
					},
					dataType: "json",
					success: (data) => {
						if (data.desc && data.desc === 'token异常，重新登陆') {
							this.$toast(data.desc);
							this.$router.push('/');
						}
						var json = data.data
						window.console.log(json)
						this.$toast(data.desc)
					}
				})
			},
			submitadd() {
				$.ajax({
					url: "http://47.98.227.139:9001/book/add",
					type: "post",
					data: {
						title: this.title,
						author: this.author,
						price: this.price,
						num: this.num,
						desc: this.desc,
						src: this.booksrc,
						token:this.token
					},
					dataType: "json",
					success: (data) => {
						if (data.desc && data.desc === 'token异常，重新登陆') {
							this.$toast(data.desc);
							this.$router.push('/');
						}
						var json = data.data
						window.console.log(json)
						this.$toast(data.desc)
					}
				})
			},
			submitdelete() {
				$.ajax({
					url: "http://47.98.227.139:9001/book/delete",
					type: "get",
					data: {
						id: this.bookid,
						title:this.title,
						token:this.token
					},
					dataType: "json",
					success: (data) => {
						if (data.desc && data.desc === 'token异常，重新登陆') {
							this.$toast(data.desc);
							this.$router.push('/');
						}
						this.$toast(data.desc)
					}
				})
			}
		},
	}
</script>

<style scoped="scoped">
	.submit {
		margin-top: 50px;
		/* border: 1px solid orangered; */
		border-radius: 5px;
		font-size: 16px;
		background-color: whitesmoke;
	}
</style>
