<template>
    <div class="postBlock">
        <div class="postTitle">
            <el-button type="text" style="font-size: 1vw; font-weight: bold"
                       @click="openCommentPage"
            >{{this.title}}</el-button>
            <p style="text-align: right; margin-top: 0">{{this.author}}</p>
        </div>
        <div class="postContent">
            <p style="text-align:left">{{this.description}}</p>
        </div>
        <div class="postFoot">
            <span style="float: left">{{this.lastUpdatedUntilNow}}</span>
            <span style="float: right">{{this.numberOfComment}} comments</span>
        </div>
    </div>
</template>

<script>
    export default {
        name: "PostBlock",
        props: {
            title: String,
            author: String,
            description: String,
            numberOfComment: Number,
            updatedTime: String,
            postId: String
        },
        computed:{
            lastUpdatedUntilNow () {
                let d = this.updatedTime;
                return d;
                // return d.getDate() + "-" + (d.getMonth() + 1) + "-" + d.getFullYear() + " " + d.getHours() + ":" + d.getMinutes()
            }
        },
        methods: {
            openCommentPage() {
                this.$router.push("/post/" + this.postId)
            }
        }
    }
</script>

<style scoped>
    .postTitle{
        height: 3vw;
        width: 100%;
        text-align: left;
    }
    .postContent{
        height: 7vw;
        width: 100%;
        text-align: left;
    }
    .postFoot{
        height: 2vw;
        width: 100%;
    }
    .postBlock{
        background-color: floralwhite;
        margin: 1vw;
        box-sizing: border-box;
        padding: 1vw;
        border-radius: 15px;
        line-height: 1vw;
    }
</style>