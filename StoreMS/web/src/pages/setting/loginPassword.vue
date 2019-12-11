<template>
	<div>
		<van-nav-bar
			title="登录密码"
			left-arrow
			@click-left="goBack"
		/>
		
		<van-cell-group>
			<van-field v-model="oldPassword" type="password" maxlength="16" placeholder="原密码" :right-icon="rightIcon" clearable/>
			<van-field v-model="newPassword1" type="password" maxlength="16" placeholder="新密码" :right-icon="rightIcon" clearable/>
			<van-field v-model="newPassword2" type="password" maxlength="16" placeholder="确认密码" :right-icon="rightIcon" clearable/>
		</van-cell-group>
		
		<van-overlay :show="show" class-name="overlay" />
		<van-loading v-if="show" vertical>{{loadingText}}</van-loading>
		
		<van-button round type="danger" @click="onSubmit">确 定</van-button>
		
	</div>
</template>

<script>
	import {NavBar,CellGroup,Field,Button,Toast,Overlay,Loading} from 'vant'
	import {Images} from '@/resources/index.js'
	import Util from '@/utils/Utils.js'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[CellGroup.name]:CellGroup,
			[Field.name]:Field,
			[Button.name]:Button,
			[Toast.name]:Toast,
			[Overlay.name]:Overlay,
			[Loading.name]:Loading
		},
		data(){
			return{
				show:false,
				loadingText:'修改中...',
				oldPassword:'',
				newPassword1:'',
				newPassword2:'',
				rightIcon:Images.login.ic_lock
			}
		},
		methods:{
			goBack(){
				this.$router.back()
			},
			onSubmit(){
				if (this.checkPwd()) {
					this.show = true
					this.loadingText = '修改中...'
					var url = '/user/pwd?phonenum='+this.$store.state.username+'&oldPwd='+this.oldPassword+'&newPwd='+this.newPassword2+'&token='+this.$store.state.token
					this.axios.get(url).then((response)=>{
						if(response.data.detail && response.data.detail === 'token异常，重新登陆') {
							Toast('登陆过期，请重新登录')
							this.$router.push('/login')
						}else if (response.data.status === 1) {
							setTimeout(()=>{
								this.loadingText = response.data.detail
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
			},
			checkPwd(){
				if (this.oldPassword === '' || !this.newPassword1 === '' || this.newPassword2 === '') {
					Toast('密码不能为空')
					return false
				}
				if (!Util.isPassword(this.newPassword1)) {
					Toast('密码必须是6位以上的字母数字符号中任意两种或三种组合')
					return false
				}
				if (this.newPassword1 !== this.newPassword2) {
					Toast('新密码不一致')
					return false
				}
				if (this.oldPassword === this.newPassword1) {
					Toast('新密码不能与原密码相同')
					return false
				}
				return true
			},
			
		}
	}
</script>

<style scoped>
	.van-nav-bar__arrow{
		color: black;
		font-size: 1.2rem;
	}
	.van-cell-group{
		margin: 0.625rem 0.625rem 0 0.625rem;
		border-radius: 0.625rem;
		overflow: hidden;
	}
	.van-button{
		width: 15rem;
		position: absolute;
		left: 0;
		right: 0;
		bottom: 0.9375rem;
		margin: 0 auto;
	}
	.overlay{
		background-color: rgba(0,0,0,0);
	}
</style>
