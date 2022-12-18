<template>
  <div class="about">
    <h1>{{ judul }}</h1>
    <input type="text" v-model="judul">
  </div>
</template>

<script>

export default {
  name: "About",
  data() {
    return {
      judul: "Rakamin - BTPN Syariah",
    }
  }
};
</script>