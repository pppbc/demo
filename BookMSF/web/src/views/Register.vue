<template>
	<div class="view">
		<!-- 顶部信息 -->
		<div class="header">
			<p>图书管理系统注册界面</p>
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
				<input type="submit" id="submit" @click="login()" value="登陆">
			</div>
		</div>

		<div class="footer">
			<p>书籍是人类进步的阶梯！</p>
		</div>
	</div>
</template>

<script>
	import $ from 'jquery'
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
		methods: {
			submit() {
				$.ajax({
					url: "http://47.98.227.139:9001/user/register",
					type: "get",
					data: {
						username: this.username,
						password: this.password
					},
					dataType: "json",
					success: (data) => {
						switch (data.status) {
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
