<template>
    <div class="primaryPage">
        <div v-for="(block, index) in this.blockResume" :key="index">
            <PostBlock :title="block.title"
                       :author="block.author"
                       :description="block.description"
                       :number-of-comment="block.numberOfComment"
                       :updated-time="block.datetime"
                       :post-id="block.postId"
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
                    withCredentials: true
                });
                let config = {
                    header: {
                        'Content-Type':'application/json'
                    },
                    params: {}
                };
                axiosAjax.get('http://localhost:8080/post', config).then((res)=>{
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