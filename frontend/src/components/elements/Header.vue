<script setup lang="ts">
    import FillerBox from '../utils/FillerBox.vue';
</script>

<template>
    <header class="bg-blue-500 text-white">
        <div class="flex flex-row gap-4 items-center p-3">
            <h1 class="text-lg text-bold">IdeaWeb</h1>
            <form class="flex flex-row relative">
                <input class="w-75 focus:outline-none bg-white text-black rounded-md p-4" type="text" placeholder="Find an author or a quote"/>
                <button class="hover:bg-blue-600 absolute right-2 top-2 rounded-md bg-blue-500 p-2" type="submit">Ctrl+K</button>
            </form>
            <FillerBox/>
            <nav class="flex flex-row gap-2">
                <a class="hover:underline" href="home">Home</a>
                <a class="hover:underline" href="profile">Profile</a>
                <a class="hover:underline" href="about">About</a>
            </nav>
        </div>
    </header>
</template>