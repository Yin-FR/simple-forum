<template>
    <div class="postBlock">
        <div class="postTitle">
            <p style="text-align:left">{{this.title}}</p>
            <p style="text-align: right; margin-top: 0">{{this.author}}</p>
        </div>
        <div class="postContent">
            <p style="text-align:left">{{this.description}}</p>
        </div>
        <div class="postFoot">
            <span style="float: left" class="header_text">
                <el-button type="info" @click="dialogVisible = true" class="commentBottom">Comment</el-button>
            </span>
            <span style="float: right">{{this.numberOfComment}} comments</span>
        </div>
        <el-dialog
                title="Comment"
                :visible.sync="dialogVisible"
                width="30%"
                :before-close="handleClose">
            <el-form ref="form" :model="form" label-width="0px">
                <el-form-item>
                    <el-input type="textarea" v-model="form.comment"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="closeComment">Cancer</el-button>
                <el-button type="primary" @click="closeComment">Confirm</el-button>
            </span>
        </el-dialog>
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
            postId: String
        },
        data() {
            return {
                dialogVisible: false,
                form: {
                    comment: ''
                }
            }
        },
        methods: {
            emptyComment() {
                this.comment = '';
            },
            closeComment(submit) {
                this.dialogVisible = false;
                if (submit) {
                    let commentSend = this.form.comment;
                    if (commentSend){
                        let data = {
                            commentContent: commentSend,
                            author: (new Date()).valueOf().toString(),
                            postId: this.postId
                        };
                        let config = {
                            header: {
                                'Content-Type':'application/json'
                            }
                        };
                        const axiosAjax = this.axios.create({
                            timeout: 1000*60,
                            withCredentials: false
                        });
                        axiosAjax.post('http://localhost:8000/comment', data, config).then((res)=>{
                            console.log(data)
                            console.log(res.data)
                            this.$notify({
                                title: "submit success",
                                type: "success",
                                message: "thanks !",
                                duration: 2000
                            });
                            this.emptyComment();
                        }).catch((err)=>{
                            console.log(err);
                            this.emptyComment();
                        });
                    } else {
                        this.$notify({
                            title: "submit error",
                            type: "error",
                            message: "error",
                            duration: 2000
                        })
                    }
                }
            },
            handleClose(done) {
                done()
                this.emptyComment();
            },
        }
    }
</script>

<style scoped>
    .postTitle{
        height: 4vw;
        width: 100%;
        text-align: left;
    }
    .postContent{
        height: 8vw;
        width: 100%;
        text-align: left;
    }
    .postBlock{
        background-color: floralwhite;
        margin: 1vw;
        box-sizing: border-box;
        padding: 1vw;
        border-radius: 15px;
        line-height: 1vw;
    }
    .postFoot{
        height: 1vw;
        width: 100%;
    }
    span.header_text {
        text-align: left;
        /*display: inline-block;*/
        float: left;
        align-items: center;
        /*padding-top: 1.2vw;*/
        box-sizing: border-box;
    }
    .commentBottom{
        position: relative;
        bottom: 1vw;
    }
</style>