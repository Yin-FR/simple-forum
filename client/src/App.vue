<template>
    <div id="app">
        <Header class="header"/>
        <PrimaryPage class="primaryPage"/>
    </div>
</template>

<script>

    import Header from "./components/Header";
    import PrimaryPage from "./views/PrimaryPage";
    export default {
        name: 'App',
        components: {
            PrimaryPage,
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

    .primaryPage{
        height: 93% !important;
    }
</style>
