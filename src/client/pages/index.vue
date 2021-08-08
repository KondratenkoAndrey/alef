<template>
  <div class="container">
    <div class="row">
      <!-- eslint-disable-next-line vue/no-v-html -->
      <div v-html="content.main"></div>
    </div>
    <div class="row text-center align-items-center justify-content-center my-4">
      <!-- prettier-ignore -->
      <div class="card bg-light m-3" style="width: 18rem"><h1>1</h1></div>
      <div class="card bg-light m-3" style="width: 18rem"><h1>2</h1></div>
      <div class="card bg-light m-3" style="width: 18rem"><h1>3</h1></div>
      <div class="card bg-light m-3" style="width: 18rem"><h1>4</h1></div>
      <div class="card bg-light m-3" style="width: 18rem"><h1>5</h1></div>
      <div class="card bg-light m-3" style="width: 18rem"><h1>6</h1></div>
      <div class="card bg-light m-3" style="width: 18rem"><h1>7</h1></div>
      <div class="card bg-light m-3" style="width: 18rem"><h1>8</h1></div>
      <div class="card bg-light m-3" style="width: 18rem"><h1>9</h1></div>
      <div class="card bg-light m-3" style="width: 18rem"><h1>10</h1></div>
      <div class="card bg-light m-3" style="width: 18rem"><h1>11</h1></div>
      <div class="card bg-light m-3" style="width: 18rem"><h1>12</h1></div>
    </div>
    <div class="row">
      <!-- eslint-disable-next-line vue/no-v-html -->
      <div v-html="content.bottomText"></div>
    </div>
  </div>
</template>

<script>
import axios from '~/plugins/axios'

export default {
  async asyncData({ route, error }) {
    const page = await axios.get('/page/' + encodeURIComponent(route.path))
    const rawContent = await axios.get('/content/' + page.data.id)
    const contents = rawContent.data.reduce(
      (a, x) => ({ ...a, [x.tag]: x.data }),
      {}
    )
    return {
      title: page.data.title,
      description: page.data.description,
      content: contents,
    }
  },
  data() {
    return {
      title: '',
      description: '',
      content: {},
    }
  },
  head() {
    return {
      title: this.title,
      meta: [
        {
          name: 'description',
          content: this.description,
        },
      ],
    }
  },
}
</script>

<style scoped>
.title {
  font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>
