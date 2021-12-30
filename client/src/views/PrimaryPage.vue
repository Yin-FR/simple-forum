<template>
    <div class="primaryPage">
        <div v-for="(post, index) in this.allPosts" :key="index">
            <PostBlock :title="post.title"
                       :author="post.author"
                       :description="post.content"
                       :number-of-comment="post.commentNumber"
                       :updated-time="'...'"
                       :post-id="post.postId"
                       class="blockClass"
            >
            </PostBlock>
        </div>
    </div>
</template>

<script>
    import PostBlock from "../components/PostBlock";
    export default {
        name: "PrimaryPage",
        components: {PostBlock},
        data() {
            return {
                blockResume: [
                    {
                        title: "Do you believe our government ?",
                        author: "tiny_Xu",
                        numberOfComment: 532,
                        datetime: new Date(Date.now()),
                        description: "syigvcgaiylgl af eliwfiagew fiuwgfilwe filwegfila egwilfgueiaugwfiluawfgliagfliwgf uilawgefliigwf",
                        postId: 1
                    },
                    {
                        title: "Do you believe our government ?",
                        author: "tiny_Xu",
                        numberOfComment: 532,
                        datetime: new Date(Date.now()),
                        description: "syigvcgaiylgl af eliwfiagew fiuwgfilwe filwegfila egwilfgueiaugwfiluawfgliagfliwgf uilawgefliigwf",
                        postId: 2
                    },
                    {
                        title: "Do you believe our government ?",
                        author: "tiny_Xu",
                        numberOfComment: 532,
                        datetime: new Date(Date.now()),
                        description: "syigvcgaiylgl af eliwfiagew fiuwgfilwe filwegfila egwilfgueiaugwfiluawfgliagfliwgf uilawgefliigwf",
                        postId: 3
                    },
                    {
                        title: "Do you believe our government ?",
                        author: "tiny_Xu",
                        numberOfComment: 532,
                        datetime: new Date(Date.now()),
                        description: "syigvcgaiylgl af eliwfiagew fiuwgfilwe filwegfila egwilfgueiaugwfiluawfgliagfliwgf uilawgefliigwf",
                        postId: 4
                    },
                    {
                        title: "Do you believe our government ?",
                        author: "tiny_Xu",
                        numberOfComment: 532,
                        datetime: new Date(Date.now()),
                        description: "syigvcgaiylgl af eliwfiagew fiuwgfilwe filwegfila egwilfgueiaugwfiluawfgliagfliwgf uilawgefliigwf",
                        postId: 5
                    },
                ],
                allPosts: []
            }
        },
        methods: {
            getAllPosts() {
                const axiosAjax = this.axios.create({
                    timeout: 60*1000,
                    withCredentials: false
                });
                let config = {
                    header: {
                        'Content-Type':'application/json'
                    },
                    params: {}
                };
                axiosAjax.get('http://localhost:8000/post', config).then((res)=>{
                    setTimeout(()=>{this.allPosts = res.data;}, 100)
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
            this.getAllPosts();
        }
    }
</script>

<style scoped>
    .primaryPage{
        background-color: lightslategray;
        padding-top: 0.5vw;
    }
    .blockClass{
        width: 50%;
        margin-left: 25%;
        box-sizing: border-box;
    }
</style>