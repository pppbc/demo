<template>
	<van-tabbar
		route
		v-model="mActive"
		@change="$emit('orderBarChange',mActive)"
		active-color="rgb(255,58,73)"
		inactive-color="#333"
		:fixed="false"
		class="orderBar"
	>
		<van-tabbar-item
			v-for="(item,index) in orderBarItems"
			:key="index"
			:name="item.name"
			:icon="item.icon"
			:to="item.to"
		>
			{{item.text}}
		</van-tabbar-item>
	</van-tabbar>
</template>

<script>
	import {Tabbar,TabbarItem} from 'vant'
	export default{
		components:{
			[Tabbar.name]:Tabbar,
			[TabbarItem.name]:TabbarItem
		},
		props:{
			orderBarItems:Array,
			active:String
		},
		data(){
			return{
				mActive:this.active
			}
		},
		watch:{
			return (val){
				this.mActive = val
			}
		}
	}
</script>

<style scoped>
	.orderBar{
		position: fixed;
		top: 2.875rem;
		left: 2.5%;
		width: 95%;
		height: 4rem;
		border-radius: 0.625rem;
	}
</style>
