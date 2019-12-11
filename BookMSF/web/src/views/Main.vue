<template>
	<div>
		<router-view></router-view>
		<foot-bar :footBarItems="footBarItems" :active="active" @footBarChange="footBarChange"></foot-bar>
	</div>
</template>

<script>
	import footBar from '../components/foot-bar.vue'
	export default {
		components: {
			footBar
		},
		data() {
			return {
				active: 'home',
				footBarItems: [{
						name: 'home',
						icon: 'wap-home',
						text: '系统首页',
						to: '/main/home',
					},
					{
						name: 'about',
						icon: 'friends',
						text: "个人中心",
						to: '/main/about',
						dot: true
					}
				]
			}
		},
		methods: {
			footBarChange(active) {
				
			}
		}
	}
</script>

<style scoped>
	body {
		margin: 0;
		padding: 0;
		background-color: #fbfaf7;
	}

	/* .head {
		position: fixed;
		top: 0;
		left: 0;
		z-index: 9;
	}

	.foot {
		width: 100%;
		position: fixed;
		bottom: 0;
		left: 0;
	} */
</style>

