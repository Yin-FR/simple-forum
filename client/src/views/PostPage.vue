<template>
    <div style="background-color: lightslategray; padding-top: 0.8vw">
        <CommentTitle
                class="commentTitle"
                :post-id="this.postDetail.postId"
                :title="this.postDetail.title"
                :author="this.postDetail.author"
                :description="this.postDetail.content"
                :number-of-comment="this.postDetail.commentNumber"
        />
        <div v-for="(comment, index) in this.postDetail.comment" :key="index">
            <CommentBlock
                    class="commentBlock"
                    :author="comment.author"
                    :description="comment.CommentContent"
            />
        </div>
    </div>
</template>

<script>
    import CommentBlock from "../components/CommentBlock";
    import CommentTitle from "../components/CommentTitle";
    export default {
        name: "PostPage",
        components: {CommentBlock, CommentTitle},
        data(){
            return {
                postDetail: {}
            }
        },
        methods: {
            getAllComments() {
                let postId = this.$route.params.postId;
                const axiosAjax = this.axios.create({
                    timeout: 60*1000,
                    withCredentials: false
                });
                let config = {
                    header: {
                        'Content-Type':'application/json'
                    },
                    params: {
                        postId: postId
                    }
                };
                axiosAjax.get('http://localhost:8000/comment', config).then((res)=>{
                    console.log(res.data)
                    setTimeout(()=>{this.postDetail = res.data;}, 100)
                }).catch((err)=>{
                    this.$notify({
                        type: 'error',
                        title: 'error',
                        message: err,
                    })
                });
            }
        },
        beforeMount() {
            this.getAllComments();
        }
    };
</script>

<style scoped>
    .commentBlock{
        width: 50%;
        margin-left: 25%;
    }
    .commentTitle{
        width: 60%;
        margin-left: 20%;
        font-weight: bold;
        box-sizing: border-box;
    }
</style>