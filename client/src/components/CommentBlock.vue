<template>
    <div class="postBlock">
        <div class="postTitle">
            <p style="text-align: right">{{this.author}}</p>
        </div>
        <div class="postContent">
            <p style="text-align:left">{{this.description}}</p>
        </div>
        <div class="postFoot">
            <span style="float: left">{{this.lastUpdatedUntilNow}}</span>
        </div>
    </div>
</template>

<script>
    export default {
        name: "PostBlock",
        props: {
            author: String,
            description: String,
            updatedTime: Date,
        },
        computed:{
            lastUpdatedUntilNow () {
                let d = this.updatedTime;
                return d.getDate() + "-" + (d.getMonth() + 1) + "-" + d.getFullYear() + " " + d.getHours() + ":" + d.getMinutes()
            }
        },
        methods: {
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