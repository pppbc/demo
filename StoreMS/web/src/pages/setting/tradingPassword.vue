<template>
	<div>
		<van-nav-bar
			title="支付密码"
			left-arrow
			@click-left="goBack"
		/>
		
		<div v-if="isFirst">
			<div v-if="setText1 === '设置新的支付密码'">
				<p style="font-size: 0.75rem; opacity: 0.7;text-align: center;">设置新的支付密码</p>
				<van-password-input
					:value="password1"
					:focused="showKeyboard"
					@focus="showKeyboard = true"
				/>
			</div>
			<div v-else>
				<p style="font-size: 0.75rem; opacity: 0.7;text-align: center;">确认新的支付密码</p>
				<van-password-input
					:value="password2"
					:focused="showKeyboard"
					@focus="showKeyboard = true"
				/>
			</div>
		</div>
		<div v-else>
			<div v-if="setText2 === '输入当前支付密码'">
				<p style="font-size: 0.75rem; opacity: 0.7;text-align: center;">输入当前支付密码</p>
				<van-password-input
					:value="password1"
					:focused="showKeyboard"
					@focus="showKeyboard = true"
				/>
			</div>
			<div v-else-if="setText2 === '输入新的支付密码'">
				<p style="font-size: 0.75rem; opacity: 0.7;text-align: center;">输入新的支付密码</p>
				<van-password-input
					:value="password2"
					:focused="showKeyboard"
					@focus="showKeyboard = true"
				/>
			</div>
			<div v-else>
				<p style="font-size: 0.75rem; opacity: 0.7;text-align: center;">确认新的支付密码</p>
				<van-password-input
					:value="password3"
					:focused="showKeyboard"
					@focus="showKeyboard = true"
				/>
			</div>
		</div>
		<van-number-keyboard
			:show="showKeyboard"
			theme="custom"
			close-button-text="完成"
			@input="onInput"
			@delete="onDelete"
			@close="onSubmit"
		/>
		
		
		<van-overlay :show="show" class-name="overlay" />
		<van-loading v-if="show" vertical>{{loadingText}}</van-loading>
	</div>
</template>

<script>
	import {NavBar,PasswordInput,NumberKeyboard,Button,Toast,Overlay,Loading} from 'vant'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[PasswordInput.name]:PasswordInput,
			[NumberKeyboard.name]:NumberKeyboard,
			[Button.name]:Button,
			[Toast.name]:Toast,
			[Overlay.name]:Overlay,
			[Loading.name]:Loading
		},
		data(){
			return{
				show:false,
				loadingText:'修改中...',
				isFirst:false,
				setText1:'输入支付密码',
				setText2:'输入原支付密码',
				password1:'',
				password2:'',
				password3:'',
				showKeyboard:true,
				
			}
		},
		methods:{
			goBack(){
				this.$router.back()
			},
			judgePaymentPW(){
				var url = '/user/getpaypwd?user_id='+this.$store.state.userID+'&token='+this.$store.state.token
				this.axios.get(url).then((response)=>{
					if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
						Toast('登陆过期，请重新登录')
						this.$router.push('/login')
					}else if (response.data.status === 1) {
						this.isFirst = false
						this.setText2 = '输入当前支付密码'
					} else{
						this.isFirst = true
						this.setText1 = '设置新的支付密码'
					}
				}).catch((response)=>{
					Toast('网络请求出错，请稍候再试!')
				})
			},
			onInput(key){
				if (this.setText1 === '设置新的支付密码') {
					this.password1 = (this.password1 + key).slice(0,6)
					if (this.password1.length === 6) {
						this.$nextTick(()=>{
							this.setText1 = '确认新的支付密码'
						})
					}
				} 
				if (this.setText1 === '确认新的支付密码') {
					this.password2 = (this.password2 + key).slice(0,6)
				}
				if (this.setText2 === '输入当前支付密码') {
					this.password1 = (this.password1 + key).slice(0,6)
					if (this.password1.length === 6) {
						this.$nextTick(()=>{
							this.setText2 = '输入新的支付密码'
						})
					}
				}
				if (this.setText2 === '输入新的支付密码') {
					this.password2 = (this.password2 + key).slice(0,6)
					if (this.password2.length === 6) {
						this.$nextTick(()=>{
							this.setText2 = '确认新的支付密码'
						})
					}
				}
				if (this.setText2 === '确认新的支付密码') {
					this.password3 = (this.password3 + key).slice(0,6)
				}
			},
			onDelete(){
				if (this.setText1 === '设置新的支付密码') {
					this.password1 = this.password1.slice(0,this.password1.length - 1)
				} 
				if (this.setText1 === '确认新的支付密码') {
					this.password2 = this.password2.slice(0,this.password2.length - 1)
				}
				if (this.setText2 === '输入当前支付密码') {
					this.password1 = this.password1.slice(0,this.password1.length - 1)
				}
				if (this.setText2 === '输入新的支付密码') {
					this.password2 = this.password2.slice(0,this.password2.length - 1)
				}
				if (this.setText2 === '确认新的支付密码') {
					this.password3 = this.password3.slice(0,this.password3.length - 1)
				}
			},
			onSubmit(){
				if (this.setText1 === '确认新的支付密码') {
					if (this.checkPwd()) {
						this.show = true
						this.loadingText = '正在设置...'
						var url = '/user/setpaypwd?user_id='+this.$store.state.userID+'&password='+this.password2+'&token='+this.$store.state.token
						this.axios.get(url).then((response)=>{
						if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
								Toast('登陆过期，请重新登录')
								this.$router.push('/login')
							}else if (response.data.status === 1) {
								setTimeout(()=>{
									this.loadingText = '设置成功'
									setTimeout(()=>{
										this.show = false
										this.$router.back()
									},500)
								},1000)
							} else{
								setTimeout(()=>{
									Toast('密码修改失败，请重试')
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
				if (this.setText2 === '确认新的支付密码') {
					if (this.checkPwd()) {
						this.show = true
						this.loadingText = '修改中...'
						var url = '/user/updatepaypwd?user_id='+this.$store.state.userID+'&oldPwd='+this.password1+'&newPwd='+this.password3+'&token='+this.$store.state.token
						this.axios.get(url).then((response)=>{
						if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
								Toast('登陆过期，请重新登录')
								this.$router.push('/login')
							}else if (response.data.status === 1) {
								setTimeout(()=>{
									this.loadingText = '修改成功'
									setTimeout(()=>{
										this.show = false
										this.$router.back()
									},500)
								},1000)
							} else{
								setTimeout(()=>{
									Toast('密码修改失败，请重试')
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
			checkPwd(){
				if (this.setText1 === '确认新的支付密码') {
					if(this.password1 !== this.password2){
						Toast('新密码不一致')
						return false
					}
					return true
				} 
				if (this.setText2 === '确认新的支付密码') {
					if (this.password2 !== this.password3) {
						Toast('新密码不一致')
						return false
					}
					if (this.password1 === this.password3) {
						Toast('新密码与原密码不能相同')
						return false
					}
					return true
				}
			}
		},
		beforeMount() {
			this.judgePaymentPW()
		}
	}
</script>

<style scoped>
	.van-nav-bar__arrow{
		color: black;
		font-size: 1.2rem;
	}
	.overlay{
		background-color: rgba(0,0,0,0);
	}
</style>
