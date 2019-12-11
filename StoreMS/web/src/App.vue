<template>
  <div id="app">
		<transition name="fade" mode="out-in">
			<router-view v-if="isRouterAlive"></router-view>
		</transition>
  </div>
</template>

<script>
	export default {
		name: 'app',
		provide(){
			return {
				reload:this.reload
			}
		},
		data(){
			return {
				isRouterAlive:true
			}
		},
		methods:{
			reload(){
				this.isRouterAlive = false
				this.$nextTick(function(){
					this.isRouterAlive = true
				})
			}
		}
	}
</script>

<style>
body{
	margin: 0;
	padding: 0;
	background: #F5F5F5;
}
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
.fade-enter{
	opacity: 0;
}
.fade-enter-active{
	transition: opacity .5s;
}
.fade-leave{
	opacity: 1;
}
.fade-leave-active{
	opacity: 0;
	transition: opacity .5s;
}
</style>
