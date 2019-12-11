<template>
	<div>
		<img class="bg" :src="bg">
		<van-nav-bar
			:title="title"
			left-arrow
			fixed
			@click-left="goBack"
		/>
		<order-bar :orderBarItems="orderBarItems" :active="title" @orderBarChange="orderBarChange"></order-bar>
		<router-view></router-view>
		
	</div>
</template>

<script>
	import {NavBar} from 'vant'
	import {Images} from '@/resources/index.js'
	import orderBar from '@/components/orderBar.vue'
	export default{
		components:{
			[NavBar.name]:NavBar,
			orderBar
		},
		data(){
			return{
				bg:Images.indent.ic_bg,
				orderBarItems:[
					{
						name:'待付款',
						icon:Images.indent.ic_tab_normal[0],
						text:'待付款',
						to:'/orderPage/unpaid/待付款'
					},
					{
						name:'待发货',
						icon:Images.indent.ic_tab_normal[1],
						text:'待发货',
						to:'/orderPage/unshipped/待发货'
					},
					{
						name:'待收货',
						icon:Images.indent.ic_tab_normal[2],
						text:'待收货',
						to:'/orderPage/unreceived/待收货'
					},
					{
						name:'待评价',
						icon:Images.indent.ic_tab_normal[3],
						text:'待评价',
						to:'/orderPage/comment/待评价'
					},
					{
						name:'全部订单',
						icon:Images.indent.ic_tab_normal[4],
						text:'全部订单',
						to:'/orderPage/allorders/全部订单'
					}
				]
			}
		},
		methods:{
			goBack(){
				this.$store.commit('changeTabbarActive','mine')
				this.$router.push('/main/mine')
			},
			orderBarChange(active){
				console.log(active)
			}
		},
		computed:{
			title(){
				return this.$route.params.name
			}	
		}
	}
</script>

<style scoped>
	.bg{
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 6.5rem;
		z-index: -1;
	}
	.van-nav-bar{
		background: url(../../assets/Images/Mine/background1.png);
	}
	.van-nav-bar__title{
		color: white;
	}
	.van-nav-bar__arrow{
		font-size: 1.5rem;
		color: white;
	}
</style>
