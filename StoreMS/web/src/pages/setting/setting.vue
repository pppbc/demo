<template>
	<div id="setting">
		<van-nav-bar
			title="设置"
			left-arrow
			@click-left="goBack"
			:border="border"
		/>
		
		<van-cell-group>
			<van-cell
				v-for="(item,index) in dataList"
				:key="index"
				:title="item.title"
				is-link
				@click="goNextPage(item.name)"
			>
			</van-cell>
		</van-cell-group>
		
		<van-button type="danger" round @click="signOut">退出当前账号</van-button>
	</div>
</template>

<script>
	import {NavBar,CellGroup,Cell,Button,Dialog} from 'vant'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[CellGroup.name]:CellGroup,
			[Cell.name]:Cell,
			[Button.name]:Button,
			[Dialog.Component.name]:Dialog.Component
		},
		data(){
			return {
				border:false,
				dataList:[
					{title:'我的收货地址',name:'/addressSetting'},
					{title:'账户与安全',name:'/accountAndSecurity'},
					{title:'音效与通知',name:'none'},
					{title:'通用',name:'none'},
					{title:'隐私',name:'none'},
					{title:'清除缓存',name:'none'},
				]
			}
		},
		methods:{
			goBack(){
				this.$store.commit('changeTabbarActive','mine')
				this.$router.back()
			},
			goNextPage(name){
				if (name !== 'none') {
					this.$router.push(name)
				}
			},
			signOut(){
				Dialog.confirm({
					title:'提示',
					message:'确定要退出吗？'
				}).then(()=>{
					this.$store.commit('token','')
					localStorage.removeItem('token')
					this.$router.push('/login')
				}).catch(()=>{
					
				})
			}
		}
	}
</script>

<style scoped>
	.van-nav-bar__arrow{
		color: black;
		font-size: 1.2rem;
	}
	.van-cell-group{
		margin: 0.625rem;
		border-radius: 0.625rem;
		overflow: hidden;
		text-align: left;
		border: 1px solid #f0f0f0;
	}
	.van-button{
		width: 15rem;
		position: absolute;
		left: 0;
		right: 0;
		bottom: 0.9375rem;
		margin: 0 auto;
	}
</style>
