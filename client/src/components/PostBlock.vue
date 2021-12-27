<template>
    <div class="postBlock">
        <div class="postTitle">
            <p>{{this.title}}</p>
            <p>{{this.author}}</p>
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
            updatedTime: Date
        },                                                                                                                                                                          
        computed:{
            lastUpdatedUntilNow () {
                let d = this.updatedTime;
                return d.getDate() + "-" + (d.getMonth() + 1) + "-" + d.getFullYear() + " " + d.getHours() + ":" + d.getMinutes()
            }
        }
    }
</script>

<style scoped>
    .postTitle{
        height: 2vw;
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
    }
</style>