<template>
	<div class="view">
		<!-- 顶部信息 -->
		<div class="header">
			<p>xx图书管理平台</p>
		</div>

		<!-- 输入框 -->
		<div class="content">
			<div class="admin">
				用户名:<input type="text" id="username" placeholder="请输入用户名" v-model="username">
			</div>
			<div class="password">
				密&nbsp;&nbsp;&nbsp;码:<input type="password" id="password" placeholder="请输入密码" v-model="password">
			</div>
			<div class="loginBtn">
				<input type="submit" id="submit" @click="submit()" value="提交">
				<input type="submit" id="submitR" @click="register()" value="注册">
				<input type="submit" id="submitL" @click="justlook()" value="游客模式">
			</div>
		</div>

		<div class="footer">
			<p>读过的书，走过的路！</p>
		</div>
	</div>
</template>

<script>
	import $ from 'jquery'
	import uuidv4 from 'uuid/v4'
	import Toast from 'vant'
	export default {
		data() {
			return {
				username: '',
				password: ''
			}
		},
		comments: {
			[Toast.name]: Toast
		},
		beforeMount: function() {
			
		},
		methods: {
			submit() {
				
				
				
				$.ajax({
					url: "http://47.98.227.139:9001/user/login",
					data: {
						username: this.username,
						password: this.password,
						udid:uuidv4()
					},
					type: "get",
					dataType: "json",
					success: (data) => {
						var json = data.data
						switch (data.status) {
							case 0:
								this.$toast(data.desc)
								break;
							case 1:
								this.$toast.success(" 欢迎你 " + $("#username").val())
								//登陆成功则存储用户名在本地
								var info = {
									username: json.username,
									id:json.id,
									age:json.age,
									head_img:json.head_img,
									tel:json.tel,
									desc:json.desc,
									token:json.token,
									root:json.is_m
								}
								info = JSON.stringify(info)
								localStorage.setItem("info", info)

								this.$router.push("/main/home")
								break;
							default:
								break;
						}
					},
					error:()=>{
						alert("can't find server!")
					}
				})
			},
			register() {
				this.$router.push({
					path: '/register'
				})
			},
			justlook() {
				this.$router.push("/justlook")
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
		padding: 50px 0 90px 0;
		font-size: 30px;
		font-weight: 900;
	}

	/* content */
	.content {
		margin: 0 20px;
		padding: 20px 0;
		height: 200px;
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
		padding: 100px 0 30px 0;
		font-size: 16px;
		font-style: oblique;
		color: gray;
	}
</style>
