<template>
    <div class="primaryPage">
        <div v-for="(block, index) in this.blockResume" :key="index">
            <PostBlock :title="block.title"
                       :author="block.author"
                       :description="block.description"
                       :number-of-comment="block.numberOfComment"
                       :updated-time="block.datetime"
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
                        description: "syigvcgaiylgl af eliwfiagew fiuwgfilwe filwegfila egwilfgueiaugwfiluawfgliagfliwgf uilawgefliigwf"
                    },
                    {
                        title: "Do you believe our government ?",
                        author: "tiny_Xu",
                        numberOfComment: 532,
                        datetime: new Date(Date.now()),
                        description: "syigvcgaiylgl af eliwfiagew fiuwgfilwe filwegfila egwilfgueiaugwfiluawfgliagfliwgf uilawgefliigwf"
                    },
                    {
                        title: "Do you believe our government ?",
                        author: "tiny_Xu",
                        numberOfComment: 532,
                        datetime: new Date(Date.now()),
                        description: "syigvcgaiylgl af eliwfiagew fiuwgfilwe filwegfila egwilfgueiaugwfiluawfgliagfliwgf uilawgefliigwf"
                    },
                    {
                        title: "Do you believe our government ?",
                        author: "tiny_Xu",
                        numberOfComment: 532,
                        datetime: new Date(Date.now()),
                        description: "syigvcgaiylgl af eliwfiagew fiuwgfilwe filwegfila egwilfgueiaugwfiluawfgliagfliwgf uilawgefliigwf"
                    },
                    {
                        title: "Do you believe our government ?",
                        author: "tiny_Xu",
                        numberOfComment: 532,
                        datetime: new Date(Date.now()),
                        description: "syigvcgaiylgl af eliwfiagew fiuwgfilwe filwegfila egwilfgueiaugwfiluawfgliagfliwgf uilawgefliigwf"
                    },
                ]
            }
        }
    }
</script>

<style scoped>
    .primaryPage{
        background-color: lightslategray;
    }
</style>