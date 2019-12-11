<template>
	<div>
		<!-- <input class="file" name="file" type="file" accept="image/*" @change="update($event)" /> -->
	</div>
</template>

<script>
	export default{
		components:{
		},
		data(){
			return{
			}
		},
		methods:{
			update(e){
				var url = '/user/upload?token=' + this.$store.state.token
				var file = e.target.files[0]
				console.log(file)
				console.log(file.name)
				let param = new FormData()
				param.append('file',file,file.name)
				param.append('user_id',this.$store.state.userID)
				console.log(param)
				this.axios.post(url,param).then((res)=>{
					console.log(res)
				}).catch((res)=>{
					console.log(res)
				})
			}
		},
	}
</script>

<style scoped>
</style>
