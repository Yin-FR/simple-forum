<template>
    <div>
        <div v-for="(comment, index) in this.allComments" :key="index">
            <CommentBlock
                    :author="comment.author"
                    :description="comment.description"
                    :updated-time="comment.dateTime"
            />
        </div>
    </div>
</template>

<script>
    import CommentBlock from "../components/CommentBlock";
    export default {
        name: "PostPage",
        components: {CommentBlock},
        data(){
            return {
                allComments: []
            }
        },
        methods: {
            getAllComments() {
                let postId = this.$route.params.postId;
                const axiosAjax = this.axios.create({
                    timeout: 60*1000,
                    withCredentials: true
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
                    setTimeout(()=>{this.allComments = res.data;}, 100)
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

</style>