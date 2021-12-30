<template>
    <div class="postBlock">
        <div class="postTitle">
            <p style="text-align: right">{{this.author}}</p>
        </div>
        <div class="postContent">
            <p style="text-align:left">{{this.description}}</p>
        </div>
    </div>
</template>

<script>
    export default {
        name: "PostBlock",
        props: {
            author: String,
            description: String
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