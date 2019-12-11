<template>
	<div id="register" :style="bg">
		<van-nav-bar
			left-arrow
			@click-left="onClickLeft"
		/>
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
					v-model="password1"
					type="password"
					:right-icon="lock"
					placeholder="请输入密码"
					@click-right-icon="showPassword"
					maxlength="16"
				/>
				<van-field
					v-model="password2"
					type="password"
					:right-icon="lock"
					placeholder="请确认密码"
					@click-right-icon="showPassword"
					maxlength="16"
				/>
				<!-- <van-row class="code-style"> 
					<van-col span="11">
						<van-field 
							input-align="center"
							placeholder="输入验证码"
							v-model="code"
						/>
					</van-col> -->
					<van-col span="10" offset="2">
						<van-button type="danger" :text="codeText" :disabled="disabled" @click="sendVerification"></van-button>
					</van-col>
				</van-row>
				<van-button type="danger" size="large" text="注 册" @click="register"></van-button>
			</van-cell-group>
		</div>
		<van-overlay :show="show" class-name="overlay" />
		<van-loading v-if="show" vertical>{{loadingText}}</van-loading>
	</div>
</template>

<script>
	import {Images} from '@/resources/index.js'
	import Util from '@/utils/Utils.js'
	import {CellGroup,Field,Button,Toast,Row,Col,NavBar,Overlay,Loading} from 'vant'
	export default{
		components:{
			[CellGroup.name]:CellGroup,
			[Field.name]:Field,
			[Button.name]:Button,
			[Toast.name]:Toast,
			[Row.name]:Row,
			[Col.name]:Col,
			[NavBar.name]:NavBar,
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
				password1:'',
				password2:'',
				verificationCode:'',
				account:Images.login.ic_account,
				lock:Images.login.ic_lock,
				isHidden:true,
				codeText:'获取验证码',
				code:'',
				loadingText:'修改中...',
				show:false,
				countDownTime:90,
				isSend:false,
				disabled:false
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
			register(){
				if (this.isPhonenum()) {
					// if (this.isPasswordEmpty() && this.isPasswordEqual() && this.isPassword() && this.isVerificationCodeEqual()) {
					 if (this.isPasswordEmpty() && this.isPasswordEqual() && this.isPassword()) {
						this.show = true
						this.loadingText = '注册中...'
						var url = '/user/register?phone='+this.username+'&password='+this.password1
						this.axios.get(url).then((response)=>{
							if (response.data.status == 1) {
								setTimeout(()=>{
									this.loadingText = '注册成功'
									setTimeout(()=>{
										this.show = false
										this.$router.back()
									},500)
								},1000)
							} else{
								setTimeout(()=>{
									Toast('注册失败，用户名已经存在!')
									this.show = false
								},1000)
							}
						}).catch((response)=>{
							setTimeout(()=>{
								Toast('网络请求出错，请稍候再试!')
								this.show = false
							},1000)
						})
					}
				}
			},
			onClickLeft(){
				this.$router.back()
			},
			sendVerification(){
				if (this.isPhonenum()) {
					this.isSend = true
					this.disabled = true
					this.inputVerification()
					this.timer = setInterval(()=>{
						if (this.countDownTime === 0) {
							this.stopCountDown()
							this.countDownTime = 90
							this.isSend = false
							this.disabled = false
						} else {
							this.codeText = `${this.countDownTime--}` + 's'
						}
					},1000)
				}
			},
			isPhonenum(){
				if (Util.isMobile(this.username)) {
					return true
				}
				Toast('请输入正确的手机号码')
				return false
			},
			isPasswordEmpty(){
				if (this.password1 || this.password2) {
					return true
				}
				Toast('密码不能为空')
				return false
			},
			isPassword(){
				if (Util.isPassword(this.password1)) {
					return true
				}
				Toast('密码必须是6位以上的字母数字符号中任意两种或三种组合')
				return false
			},
			isPasswordEqual(){
				if (this.password1 === this.password2) {
					return true
				}
				Toast('两次输入的密码不一致')
				return false
			},
			isVerificationCodeEqual(){
				if (this.verificationCode === this.code) {
					return true
				}
				Toast('请输入正确的验证码')
				return false
			},
			inputVerification(){
				this.axios.get('http://94.191.12.62:8090/supply/sendMes?phone='+this.username).then((response)=>{
					if (response.data.status === 1) {
						this.code = response.data.code
					} else{
						Toast('网络请求超时，请稍候再试！')
					}
				}).catch((response)=>{
					Toast("网络请求出错，请稍候再试！")
				})
			},
			stopCountDown(){
				clearInterval(this.timer)
			}
		},
		watch:{
			
		}
	}
</script>

<style scoped>
	#register{
		width: 100%;
		height: 100%;
		position: fixed;
	}
	.van-nav-bar{
		position: absolute;
		background: rgba(0,0,0,0);
	}
	.van-nav-bar__arrow{
		color: #fff;
		font-size: 1.5rem;
	}
	.content{
		width: 80%;
		height: 25rem;
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
		margin-bottom: 2.3125rem;
	}
	.van-cell-group{
		margin: 0 1.875rem;
	}
	.van-button{
		margin-top: 1.25rem;
		font-size: 1.125rem;
		height: 2.5rem;
		line-height: 2.5rem;
	}
	.code-style{
		margin-top: 1.25rem;
	}
	.code-style .van-field{
		border: 1px solid #FF0000;
		border-radius: 0.25rem;
		height: 2.5rem;
		margin: 0;
	}
	.code-style .van-button{
		font-size: 0.75rem;
		margin: 0;
		width: 100%;
	}
	.overlay{
		background-color: rgba(0,0,0,0);
	}
</style>
