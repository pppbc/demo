<template>
	<van-tabbar
		route
		v-model="mActive"
		@change="$emit('footBarChange',mActive)"
		active-color="#FF500A"
		inactive-color="#333"
		class="footBar"
	>
		<van-tabbar-item
			v-for="(footBar,index) in footBarItems"
			:key="index"
			:name="footBar.name"
			:icon="mActive === footBar.name ? footBar.icon[1] : footBar.icon[0]"
			:to="footBar.to"
		>
			{{footBar.text}}
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
			footBarItems:Array,
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
	.footBar{
		background-color: white;
	}
	@supports(bottom:constant(safe-area-inset-bottom)) {
		.footBar{
			padding-bottom: constant(safe-area-inset-bottom);
		}
	}
</style>
