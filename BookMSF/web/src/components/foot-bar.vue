<template>
	<div>
		<!-- v-model默认绑定选中标签的索引值，通过修改v-model即可切换选中的标签 -->
		<!-- change切换标签时触发 -->
		<van-tabbar v-model="mActive" @change="$emit('footBarChange',mActive)" active-color="mediumslateblue" inactive-color="#333"
		 class="footBar" swipeable>
			<!-- name标签名称，作为匹配的标识符 -->
			<!-- to路由跳转对象，同vue-router的to属性 -->
			<van-tabbar-item v-for="(footBar,index) in footBarItems" :key="index" :name="footBar.name" :icon="footBar.icon" :to="footBar.to"
			 :dot="footBar.dot">
				{{footBar.text}}
			</van-tabbar-item>
		</van-tabbar>
	</div>
</template>

<script>
	// 导入Tabbar标签栏组件
	import {
		Tabbar,
		TabbarItem
	} from 'vant'
	export default {
		components: {
			[Tabbar.name]: Tabbar,
			[TabbarItem.name]: TabbarItem,
		},
		props: {
			footBarItems: Array,
			active: String
		},
		data() {
			// window.console.log('this.active:'+this.active)
			return {
				mActive: this.active
			}
		},
		watch: {
			return (val) {
				this.mActive = val
			}
		},
		beforeUpdate:function(){
			
		}
	}
</script>

<style scoped="scoped">
	.footBar {
		background-color: #f7f8fa;
	}
</style>
