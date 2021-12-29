<template>
    <div id="app">
        <Header class="header"/>
        <router-view></router-view>
    </div>
</template>

<script>

    import Header from "./components/Header";
    export default {
        name: 'App',
        components: {
            Header,
        }
    }

</script>

<style>

    #app {
        font-family: Avenir, Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
        margin: 0 auto;
        height: 100%;
    }

    .header{
        height: 7% !important;
    }

</style>
