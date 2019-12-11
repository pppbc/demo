<template>
	<div id="mall">
		<img class="bg" :src="bg" alt="">
		<van-nav-bar 
			title="活动"
			:border="border"
			fixed
		/>
		
		<my-integral :avatar="avatar" :allBlock="allBlock"></my-integral>
		
		<van-list
			v-model="loading"
			:finished="finished"
			finished-text="没有更多了"
			@load="onLoad"
		>
			<van-image v-for="(item,index) in allActivities" :key="index" :src="item.Image"/>
		</van-list>
	</div>
</template>

<script>
	import bus from '@/utils/bus.js'
	import {NavBar,List,Image} from 'vant'
	import {Images} from '../resources/index.js'
	import myIntegral from '@/components/myIntegral.vue'
	export default{
		components:{
			[NavBar.name]:NavBar,
			[List.name]:List,
			[Image.name]:Image,
			myIntegral,
		},
		data(){
			return{
				bg:Images.common.ic_header_bg,
				border:false,
				loading:false,
				finished:false
			}
		},
		methods:{
			onLoad(){
				setTimeout(()=>{
					bus.$emit('getAllActivities','获取所有活动页')
					this.loading = false
					this.finished = true
				},1000)
			}
		},
		computed:{
			avatar(){
				return this.$store.state.avatar
			},
			allBlock(){
				return this.$store.getters.allBlock
			},
			allActivities(){
				return this.$store.state.allActicities
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
	.van-list{
		margin: 0.625rem 0.625rem 0 0.625rem;
	}
	.van-list .van-image{
		width: 100%;
		height: auto;
		border-radius: 0.3125rem;
		overflow: hidden;
		margin-bottom: 0.625rem;
		
	}
</style>
