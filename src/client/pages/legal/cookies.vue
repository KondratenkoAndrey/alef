<template>
  <div class="container">
    <div class="row">
      <!-- eslint-disable-next-line vue/no-v-html -->
      <div v-html="contents.main"></div>
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
      contents,
    }
  },
  data() {
    return {
      title: '',
      description: '',
      contents: {},
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
