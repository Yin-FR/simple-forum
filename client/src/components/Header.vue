<template>
    <div class="header">
        <img src="../assets/IMT_Atlantique_logo_RVB.svg" alt="logo"/>
        <span>
            Simple Forum
        </span>
    </div>
</template>

<script>
    export default {
        name: "Header"
    }
</script>

<style scoped>
    .header{
        height: 100%;
        width: 100%;
        align-items:center;
        display:-webkit-flex;
    }

    img{
        height: 100%;
        width: 10vw;
        float: left;
    }

    span{
        text-align: left;
        display: inline-block;
        float: left;
    }

</style>