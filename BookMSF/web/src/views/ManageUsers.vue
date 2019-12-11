<template>
	<div>
		<van-nav-bar title="用户管理中心" fixed left-text="返回" left-arrow @click-left="onClickLeft" />
		<div style="margin-top:45px"></div>
		<!-- 功能版块 -->
		<grid :gridItems="gridItems" :columnNum="columnNum" :showClass="showClass" :gutter="gutter"></grid>

		<!-- 修改用户权限 -->
		<van-popup v-model="$store.state.showChangeU" position="bottom" :style="{height:'80%'}" @click-overlay="clickoverlay">
			<div style="margin:20px 10px;color:gray;">修改用户权限</div>
			<van-search v-model="namesearch" placeholder="搜索用户名" show-action @search="onSearch">
				<!-- slot="action"自定义插槽 -->
				<van-icon slot="action" @click="onSearch">搜索</van-icon>
			</van-search>
			<van-cell-group>
				<van-field class="ism" v-model="ism" type="number" label="权限" placeholder="请输入权限" />
			</van-cell-group>
			<van-cell-group>
				<van-field v-model="name" label="姓名" placeholder="请输入姓名" disabled />
				<van-field v-model="id" label="ID" placeholder="请输入ID" disabled />
				<van-field v-model="tel" type="number" label="联系电话" placeholder="请输入联系电话" disabled />
				<van-field v-model="age" type="number" label="年龄" placeholder="请输入年龄" disabled />
				<van-field v-model="desc" label="描述" type="textarea" placeholder="请输入个人描述" rows="2" autosize disabled />
			</van-cell-group>
			<input class="submit" @click="submitchange()" type="submit" value="确认修改">
		</van-popup>

		<!-- 删除用户 -->
		<van-popup v-model="$store.state.showDeleteU" position="bottom" :style="{height:'80%'}" @click-overlay="clickoverlay">
			<div style="margin:20px 10px;color:gray;">删除用户信息</div>
			<van-search v-model="namesearch" placeholder="搜索用户名" show-action @search="onSearch">
				<van-icon slot="action" @click="onSearch">搜索</van-icon>
			</van-search>

			<van-cell-group>
				<van-field v-model="name" label="姓名" placeholder="请输入姓名" disabled />
				<van-field v-model="id" label="ID" placeholder="请输入ID" disabled />
				<van-field class="ism" v-model="ism" type="number" label="权限" placeholder="请输入权限" disabled />
				<van-field v-model="tel" type="number" label="联系电话" placeholder="请输入联系电话" disabled />
				<van-field v-model="age" type="number" label="年龄" placeholder="请输入年龄" disabled />
				<van-field v-model="desc" label="描述" type="textarea" placeholder="请输入个人描述" rows="2" autosize disabled />
			</van-cell-group>
			<input class="submit" @click="submitdelete()" type="submit" value="确认删除">
		</van-popup>

		<!-- 搜索用户 -->
		<van-popup v-model="$store.state.showSearchU" position="bottom" :style="{height:'80%'}" @click-overlay="clickoverlay">
			<div style="margin:20px 10px;color:gray;">搜索用户信息</div>
			<van-search v-model="namesearch" placeholder="搜索用户名" show-action @search="onSearch">
				<van-icon slot="action" @click="onSearch">搜索</van-icon>
			</van-search>

			<van-cell-group>
				<van-field v-model="name" label="姓名" placeholder="请输入姓名" disabled />
				<van-field v-model="id" label="ID" placeholder="请输入ID" disabled />
				<van-field class="ism" v-model="ism" type="number" label="权限" placeholder="请输入权限" disabled />
				<van-field v-model="tel" type="number" label="联系电话" placeholder="请输入联系电话" disabled />
				<van-field v-model="age" type="number" label="年龄" placeholder="请输入年龄" disabled />
				<van-field v-model="desc" label="描述" type="textarea" placeholder="请输入个人描述" rows="2" autosize disabled />
			</van-cell-group>
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
				gridItems: [{
						text: '修改用户权限',
						path: '',
						icon: 'info-o',
						method: 'changeu'
					},
					{
						text: '删除用户',
						path: '',
						icon: 'close',
						method: 'deleteu'
					},
					{
						text: '搜索用户',
						path: '',
						icon: 'search',
						method: 'searchu'
					},
					{
						text: '展示所有信息',
						path: '/allusers',
						icon: 'description'
					}
				],
				namesearch: '',
				columnNum: 4,
				showClass: 'gridManage',
				gutter: 0,
				name: '用户名',
				id: 0,
				age: 0,
				tel: 0,
				ism: 0,
				lend: '',
				desc: '这个人很懒，什么都没有说',
				token:''
			}
		},
		beforeMount:function(){
			var info = localStorage.getItem("info")
			info = JSON.parse(info)
			this.token=info.token	
		},
		methods: {
			onClickLeft() {
				this.$router.go(-1)
			},
			onSearch() {
				$.ajax({
					url: "http://47.98.227.139:9001/user/getname",
					type: "get",
					data: {
						username: this.namesearch,
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
							this.id = json.id
							this.age = json.age
							this.tel = json.tel
							this.ism = json.is_m
							this.desc = json.desc
							this.lend = ''
							this.name = json.username
						}
						window.console.log(json)
						this.$toast(data.desc)
					},
					error: function() {
						alert("can't find server")
					}
				})
			},
			clickoverlay() {
				this.$store.state.showChangeU = false
				this.$store.state.showDeleteU = false
				this.$store.state.showSearchU = false
				this.namesearch = ''
				this.id = 0
				this.age = 0
				this.tel = ''
				this.ism = 0
				this.lend = ''
				this.desc = '这个人很懒，什么都没有说'
			},
			submitchange() {
				$.ajax({
					url: "http://47.98.227.139:9001/user/changea",
					type: "post",
					data: {
						is_m: this.ism,
						username: this.namesearch,
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
					url: "http://47.98.227.139:9001/user/delete",
					type: "get",
					data: {
						username: this.namesearch,
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
		}
	}
</script>

<style scoped="scoped">
	.ism {
		line-height: 50px;
	}

	.submit {
		margin-top: 30px;
		border-radius: 10px;
		font-size: 16px;
		background-color: whitesmoke;
	}
</style>
