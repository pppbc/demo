<template>
	<van-swipe :autoplay="3000">
		<van-swipe-item v-for="(item,index) in swipeItems" :key="index">
			<img v-lazy="item.Image">
		</van-swipe-item>
	</van-swipe>
</template>

<script>
	import {Swipe,SwipeItem} from 'vant'
	export default{
		components:{
			[Swipe.name]:Swipe,
			[SwipeItem.name]:SwipeItem
		},
		props:{
			swipeItems:Array
		}
	}
</script>

<style scoped>
	.van-swipe{
		margin: 0.625rem 0.625rem 0 0.625rem;
		border-radius: 0.3125rem;
	}
	img{
		width: 100%;
		height: 10rem;
		display: block;
	}
</style>
