<template>
	<div id="modifyInfo">
		<van-nav-bar
			:title="title"
			left-arrow
			right-text='确定'
			@click-left="goBack"
			@click-right="submit"
		/>
		
		<van-field 
			v-if="title === '昵称'" 
			placeholder="请设置昵称"
			v-model="nickname"
			clearable
		>
		</van-field>
		
		<van-field
			v-if="title === '邮箱'" 
			placeholder="请设置邮箱"
			v-model="email"
			clearable
		>
		</van-field>
		
		<van-radio-group v-if="title === '性别'" v-model="gender">
			<van-cell-group>
				<van-cell title="男" clickable @click="gender == '男'">
					<van-radio slot="right-icon" name='男'/>
				</van-cell>
				<van-cell title="女" clickable @click="gender == '女'">
					<van-radio slot="right-icon" name='女'/>
				</van-cell>
				<van-cell title="保密" clickable @click="gender == '保密'">
					<van-radio slot="right-icon" name='保密'/>
				</van-cell>
			</van-cell-group>
		</van-radio-group>
		
		<van-overlay :show="show" class-name="overlay" />
		<van-loading v-if="show" vertical>{{loadingText}}</van-loading>
	</div>
</template>

<script>
	import util from '@/utils/Utils.js'
	import {NavBar,Field,CellGroup,Cell,RadioGroup,Radio,Toast,Overlay,Loading} from 'vant'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Field.name]:Field,
			[CellGroup.name]:CellGroup,
			[Cell.name]:Cell,
			[RadioGroup.name]:RadioGroup,
			[Radio.name]:Radio,
			[Toast.name]:Toast,
			[Overlay.name]:Overlay,
			[Loading.name]:Loading
		},
		data(){
			return{
				title:'',
				placeholder:'',
				nickname:'',
				gender:'',
				email:'',
				loadingText:'登录中...',
				show:false
			}
		},
		methods:{
			goBack(){
				this.$router.back()
			},
			submit(){
				if (this.title === '昵称') {
					if (this.nickname === '') {
						Toast('昵称不能为空')
						return
					}
					this.show = true
					this.loadingText = '修改中...'
					var url = '/user/nick?user_id='+this.$store.state.userID+'&nick='+this.nickname+'&token='+this.$store.state.token
					this.axios.get(url).then((response)=>{
						if (response.data.detail && response.data.detail == 'token异常，重新登陆') {
							Toast('登陆过期，请重新登录')
							this.$router.push('/login')
						}else if (response.data.status == 1) {
							setTimeout(()=>{
								this.loadingText = response.data.detail
								this.$store.commit('nickname',this.nickname)
								setTimeout(()=>{
									this.show = false
									this.$router.back()
								},500)
							},1000)
						} else{
							setTimeout(()=>{
								Toast(response.data.detail)
								this.show = false
							},1000)
						}
					}).catch((response)=>{
						setTimeout(()=>{
							Toast('网络请求出错，请稍候再试!')
							this.show = false
						},1000)
					})
				} else if (this.title === '性别') {
					if (this.gender === '') {
						Toast('请选择您的性别')
						return
					}
					this.show = true
					this.loadingText = '修改中...'
					var url = '/user/sex?user_id='+this.$store.state.userID+'&sex='+this.gender+'&token='+this.$store.state.token
					this.axios.get(url).then((response)=>{
						if (response.data.detail && response.data.detail == 'token异常，重新登陆') {
							Toast('登陆过期，请重新登录')
							this.$router.push('/login')
						}else if (response.data.status == 1) {
							setTimeout(()=>{
								this.loadingText = response.data.detail
								this.$store.commit('gender',this.gender)
								setTimeout(()=>{
									this.show = false
									this.$router.back()
								},500)
							},1000)
						} else{
							setTimeout(()=>{
								Toast(response.data.detail)
								this.show = false
							},1000)
						}
					}).catch((response)=>{
						setTimeout(()=>{
							Toast('网络请求出错，请稍候再试!')
							this.show = false
						},1000)
					})
				} else{
					if (this.email === '') {
						Toast('邮箱不能为空')
						return 
					}
					this.show = true
					this.loadingText = '修改中...'
					var url = '/user/email?user_id='+this.$store.state.userID+'&email='+this.email+'&token='+this.$store.state.token
					this.axios.get(url).then((response)=>{
						if (response.data.detail && response.data.detail == 'token异常，重新登陆') {
							Toast('登陆过期，请重新登录')
							this.$router.push('/login')
						}else if (response.data.status == 1) {
							setTimeout(()=>{
								this.loadingText = response.data.detail
								this.$store.commit('email',this.email)
								setTimeout(()=>{
									this.show = false
									this.$router.back()
								},500)
							},1000)
						} else{
							setTimeout(()=>{
								Toast(response.data.detail)
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
		mounted(){
			this.nickname = this.$store.state.nickname ?  this.$store.state.nickname : ''
			this.gender = this.$store.state.gender ? this.$store.state.gender : ''
			this.email = this.$store.state.email ? this.$store.state.email : ''
		},
		beforeRouteEnter(to,from,next){
			next(vm=>{
				vm.title = to.meta.title
			})
		}
	}
</script>

<style scoped>
	.van-nav-bar__arrow{
		color: black;
		font-size: 1.2rem;
	}
	.van-nav-bar__text{
		color: rgb(96,96,96);
	}
	.van-radio-group{
		text-align: left;
	}
	.overlay{
		background-color: rgba(0,0,0,0);
	}
</style>
