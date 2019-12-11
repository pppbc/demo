<template>
	<div class="view">
		<!-- 顶部信息 -->
		<div class="header">
			<p>修改密码</p>
		</div>

		<!-- 输入框 -->
		<div class="content">
			<div class="admin">
				用户名:<input type="text" id="admin" placeholder="请输入用户名" v-model="admin">
			</div>
			<div class="password">
				密&nbsp;&nbsp;&nbsp;码:<input type="password" id="password" placeholder="请输入密码" v-model="password">
			</div>
			<div class="password">
				新密码:<input type="password" id="newpassword1" placeholder="请输入密码" v-model="newpassword1">
			</div>
			<div class="password">
				请确认:<input type="password" id="newpassword2" placeholder="请再输一次" v-model="newpassword2">
			</div>
			<div class="loginBtn">
				<input type="submit" id="submit" @click="submit()" value="提交">
				<input type="submit" id="submit" @click="login()" value="登陆">
			</div>
		</div>

		<div class="footer">
			<p>footer</p>
		</div>
	</div>
</template>

<script>
	import {
		Toast
	} from 'vant'
	import $ from 'jquery'
	
	export default {
		data() {
			return {
				admin: '',
				password: '',
				newpassword1: '',
				newpassword2:'',
				token:''
			}
		},
		comments: {
			[Toast.name]: Toast
		},
		beforeMount: function() {
			var info = localStorage.getItem("info")
			info = JSON.parse(info)
			this.admin=info.userName
			this.token=info.token
		},
		methods: {
			submit() {
				var url = "http://47.98.227.139:9001/user/changepassword"
				$.ajax({
					url: url,
					type: "post",
					data:{
						username:this.admin,
						password:this.password,
						newpassword1:this.newpassword1,
						newpassword2:this.newpassword2,
						token:this.token
					},
					dataType: "json",
					success: (data)=> {
						switch (data.code) {
							case 0:
								this.$toast(data.desc)
								break;
							case 1:
								this.$toast(data.desc)
								this.$router.push({
									path: '/'
								})
								break;
							default:
								break;
						}
					},
					error:()=>{
						alert("can't find server")
					}
				})
			},
			login() {
				this.$router.push({
					path: '/'
				})
			}

		}

	}
</script>

<style scoped="scoped">
	body {}

	input {
		border: 2px solid gainsboro;
	}

	.view {
		padding: 0;
		margin: 0;
		background-image: url(../assets/images/timg.jpg);
	}

	/* header */
	.header {
		padding: 50px 0 50px 0;
		font-size: 30px;
		font-weight: 900;
	}

	/* content */
	.content {
		margin: 0 20px;
		padding: 20px 0;
		height: 270px;
		background-color: rgba(1, 1, 1, 0.1);
	}

	.admin,
	.password,
	.loginBtn {
		margin: 20px 0;
	}

	.admin>input,
	.password>input {
		width: 210px;
		height: 40px;
	}

	.loginBtn>input {
		margin: 10px;
	}

	/* footer */
	.footer {
		/* margin: 100px 0 0 0; */
		padding: 70px 0 30px 0;
		font-size: 16px;
		font-style: oblique;
		color: gray;
	}
</style>
