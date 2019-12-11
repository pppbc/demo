<template>
	<div id="mall">
		<img class="bg" :src="bg" alt="">
		<van-nav-bar 
			title="商家"
			:border="border"
			fixed
		/>
		
		<my-integral :avatar="avatar" :allBlock="allBlock"></my-integral>
		
		<swipe :swipeItems="allActicities"></swipe>
		
		<div v-if="businessArrayIsEmpty" style="text-align: center; margin-top: 0.625rem;">
			<van-row class="content">
				<van-col span="8" v-for="(item,index) in businessArray" :key="index" >
					<van-row class="scrollContent">
						<van-col span="24" :style="backgroundImg(Math.floor(Math.random()*4))">
							<van-row>
								<van-col span="24">{{item.PMName}}</van-col>
							</van-row>
							<van-row>
								<van-col span="24"><span>{{item.balance}}</span>积分</van-col>
							</van-row>
							<van-row>
								<van-col span="24">
									<van-button @click="receivePoints(item)"></van-button>
								</van-col>
							</van-row>
						</van-col>
					</van-row>
				</van-col>
			</van-row>
		</div>
		
		<div v-else style="text-align: center;margin-top: 0.625rem; font-size: 1.125rem; color: #656565;">暂无可领取积分</div>
		
	</div>
</template>

<script>
	import bus from '@/utils/bus.js'
	import {NavBar,Toast,Row,Col,Button} from 'vant'
	import {Images} from '../resources/index.js'
	import myIntegral from '@/components/myIntegral.vue'
	import swipe from '@/components/swipe.vue'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[Toast.name]:Toast,
			[Row.name]:Row,
			[Col.name]:Col,
			[Button.name]:Button,
			myIntegral,
			swipe
		},
		data(){
			return{
				bg:Images.common.ic_header_bg,
				border:false,
			}
		},
		methods:{
			receivePoints(item){
				bus.$emit('receivePoints',item)
			}
		},
		computed:{
			avatar(){
				return this.$store.state.avatar
			},
			allActicities(){
				return this.$store.state.allActicities.slice(0,5)
			},
			allBlock(){
				return this.$store.getters.allBlock
			},
			businessArray(){
				return this.$store.state.businessArray
			},
			businessArrayIsEmpty(){
				return this.$store.state.businessArray.length > 0 ? true : false
			},
			backgroundImg(){
				return function(index){
					return{
						backgroundImage:"url("+Images.home.ic_integral_data[index]+")",
						backgroundRepeat:"no-repeat",
						backgroundSize:"100% 100%",
					}
				}
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
		background: url(../assets/Images/Home/header_bg.png);
	}
	.van-nav-bar__title{
		color: #fff;
		font-size: 1.125rem;
	}
	.van-grid{
		margin-top: 0.625rem;
	}
	.content{
		margin: 0 0.3125rem 0 0.3125rem;
		padding: 0;
	}
	.scrollContent{
		margin: 0.2rem;
		color: white;
	}
	.scrollContent .van-button{
		width: 3.125rem;
		height: 1.125rem;
		border: none;
		background: url('../assets/Images/Home/input.png') no-repeat center;
		background-size: 100% 100%;	
	}
</style>
