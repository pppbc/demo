<template>
	<!-- 列数：column-num -->
	<van-grid :column-num="columnNum" clickable :gutter="gutter">
		<van-grid-item v-for="(grid,index) in gridItems" :class="grid.show" :id="showClass" :key="index" :icon="grid.icon"
		 :text="grid.text" :to="grid.path" @click="allMethods(grid.method)" class="mouseOver" />
	</van-grid>
</template>

<script>
	import $ from 'jquery'
	import {
		Grid,
		GridItem,
		Toast
	} from 'vant'
	export default {
		data() {
			return {
			}
		},
		props: {
			gridItems: Array,
			columnNum: Number,
			showClass: String,
			gutter: Number
		},
		components: {
			[Grid.name]: Grid,
			[GridItem.name]: GridItem,
			[Toast.name]: Toast
		},
		methods: {
			allMethods(info) {
				if (info == "add") {
					this.$store.state.showAddB = true
				}
				if (info == "change") {
					this.$store.state.showChangeB = true
				}
				if (info == "delete") {
					this.$store.state.showDeleteB = true
				}
				if (info == "search") {
					this.$store.state.showSearchB = true
				}
				if (info == "changeu") {
					this.$store.state.showChangeU = true
				}
				if (info == "searchu") {
					this.$store.state.showSearchU = true
				}
				if (info == "deleteu") {
					this.$store.state.showDeleteU = true
				}
				if (info == "lendbook") {
					this.$store.state.lendbook = true
				}
				if (info == "returnbook") {
					this.$store.state.returnbook = true
				}
				if (info == "showlist") {
					this.$store.state.showlist = true
					var info = localStorage.getItem("info")
					info = JSON.parse(info)
					$.ajax({
						url: "http://47.98.227.139:9001/lend/getall",
						type: "get",
						data:{
							token:info.token
						},
						dataType: "json",
						success: (data) => {
							var json = data.data
							var html=''
							window.console.log(json)
							$.each(json, function(index, item) {
								html += "<tr>";
								html += "<td>" + item.username + "</td>";
								html += "<td>" + item.bookname + "</td>";
								html += "<td>" + item.lend_time + "</td>";
								html += "<td>" + item.is_r + "</td>";
								html += "<td>" + item.return_time + "</td>";
								html += "</tr>";
							})	
							$("#allinfo").html(html)
							html=''
						},
						error: function() {
							alert("can't find server")
						}
					})
				}
			},
			clickGrid(event) {

			}
		}
	}
</script>

<style>
	.IsShow {
		display: none;
	}

	#gridHome {
		margin: 5px 20px;
		border: 2px dotted mediumslateblue;
		color: mediumslateblue;
	}

	#gridManage {
		margin: 10px 15px;
		border: 2px dotted mediumturquoise;
		color: mediumturquoise;
	}
	#gridManage:hover{
		border:	2px dotted lightcoral;
	}
</style>
