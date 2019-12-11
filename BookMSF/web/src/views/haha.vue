<template>
	<div>
		<van-nav-bar title="错误!!" fixed left-text="返回" left-arrow @click-left="onClickLeft" />
		<div style="margin-top:200px">你看见了不该看的</div>
	</div>
</template>

<script>
	export default {
		methods: {
			onClickLeft() {
				this.$router.go(-1)
			}
		}
	}
</script>

<style>
</style>
