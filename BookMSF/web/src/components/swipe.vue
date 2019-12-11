<template>
	<!-- autoplay:自动轮播间隔 -->
	<!-- indicator-color:指示器颜色 -->
	<van-swipe :autoplay="autoplay" indicator-color="orange">
		<van-swipe-item v-for="(image,index) in swipeItems" :key="index">
			<img v-lazy="image"/>
		</van-swipe-item>
	</van-swipe>
</template>

<script>
	import {
		Swipe,
		SwipeItem
	} from 'vant'
	export default {
		components: {
			[Swipe.name]: Swipe,
			[SwipeItem.name]: SwipeItem
		},
		props: {
			swipeItems: Array,
			autoplay:Number
		}
	}
</script>

<style>
</style>
