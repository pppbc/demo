<template>
	<div id="login" :style="bg">
		<div class="content">
			<div class="logo">
				<img src="../assets/Images/Login/logo.png" alt="logo">
			</div>
			<van-cell-group>
				<van-field 
					v-model="username"
					:right-icon="account"
					placeholder="请输入手机号"
				/>
				<van-field 
					v-model="password"
					type="password"
					:right-icon="lock"
					placeholder="请输入密码"
					@click-right-icon="showPassword"
				/>
				<van-button type="danger" size="large" text="登 录" @click="login"></van-button>
				<van-row>
					<van-col span="7" @click="retrievePassword">忘记密码?</van-col>
					<van-col span="5" offset="12" @click="register">注册</van-col>
				</van-row>
			</van-cell-group>
		</div>
		<van-overlay :show="show" class-name="overlay" />
		<van-loading v-if="show" vertical>{{loadingText}}</van-loading>
	</div>
</template>

<script>
	import {Images} from '@/resources/index.js'
	import Util from '@/utils/Utils.js'
	import uuidv4 from 'uuid/v4'
	import {CellGroup,Field,Button,Toast,Row,Col,Overlay,Loading} from 'vant'
	export default{
		components:{
			[CellGroup.name]:CellGroup,
			[Field.name]:Field,
			[Button.name]:Button,
			[Toast.name]:Toast,
			[Row.name]:Row,
			[Col.name]:Col,
			[Overlay.name]:Overlay,
			[Loading.name]:Loading
		},
		data(){
			return{
				bg:{
					backgroundImage:"url("+Images.login.ic_bg+")",
					backgroundRepeat:"no-repeat",
					backgroundSize:"100% 100%",
				},
				username:'',
				password:'',
				account:Images.login.ic_account,
				lock:Images.login.ic_lock,
				isHidden:true,
				loadingText:'登录中...',
				show:false
			}
		},
		methods:{
			showPassword(){
				if (this.isHidden) {
					this.isHidden = false
					this.lock = Images.login.ic_unlock
				} else{
					this.isHidden = true
					this.lock = Images.login.ic_lock
				}
			},
			login(){
				if(this.verification()){
					this.show = true
					this.loadingText = '登录中...'
					// let param = new URLSearchParams()
					// param.append('phone', this.username)
					// param.append('password', this.password)
					// param.append('udid', uuidv4())
					// var url='/user/login'
					var url = '/user/login?phone='+this.username+'&password='+this.password+'&udid='+uuidv4()
					this.axios.get(url).then((response)=>{
						window.console.log(response)
						if (response.data.status === 1){
							setTimeout(()=>{
								this.loadingText = '登录成功'
								this.$store.commit('username',this.username)
								this.$store.commit('token',response.data.info)
								localStorage.setItem("username",this.username)
								localStorage.setItem('token',response.data.info)
								setTimeout(()=>{
									this.show = false
									localStorage.setItem("isOnce",false)
									this.$router.push('/main/home')
								},500)
							},1000)
						} else {
							setTimeout(()=>{
								Toast("用户名或密码错误！")
								this.show = false
							},1000)
						}
					}).catch((response)=>{
						setTimeout(()=>{
							Toast("网络请求出错，请稍候再试！")
							this.show = false
						},1000)
					})
				}
			},
			retrievePassword(){
				this.$router.push('/retrievePassword')
			},
			register(){
				this.$router.push('/register')
			},
			verification(){
				if (Util.isMobile(this.username)) {
					return true
				}
				Toast('请输入正确的手机号码')
				return false
			}
		}
	}
</script>

<style scoped>
	#login{
		width: 100%;
		height: 100%;
		position: fixed;
	}
	.content{
		width: 80%;
		height: 23rem;
		background: #fff;
		border-radius: 10px;
		margin: 8.75rem auto 0;
	}
	.logo{
		width: 3.125rem;
		margin: 0 auto;
	}
	.logo img{
		width: 100%;
		height: auto;
		margin-top: 1.9375rem;
		margin-bottom: 3.25rem;
	}
	.van-cell-group{
		margin: 0 1.875rem;
	}
	.van-button{
		margin-top: 2.8125rem;
		font-size: 1.125rem;
		height: 2.5rem;
		line-height: 2.5rem;
	}
	.van-row{
		margin-top: 0.625rem;
		font-size: 0.75rem;
	}
	.van-col:nth-child(1){
		color: #808080;
	}
	.van-col:nth-child(2){
		color: red;
	}
	.overlay{
		background-color: rgba(0,0,0,0);
	}
</style>
