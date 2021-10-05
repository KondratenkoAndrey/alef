<template>
  <div class="container">
    <div class="row">
      <!-- eslint-disable-next-line vue/no-v-html -->
      <div v-html="contents.main"></div>
    </div>
    <div>
      <b-card-group columns>
        <b-card v-for="service of services" :key="service.id">
          <template #header>
            <h5 class="mb-0 text-center">{{ service.name }}</h5>
          </template>
          <b-card-text>
            <picture>
              <source :srcset="service.imagePath + '.webp'" type="image/webp" />
              <source :srcset="service.imagePath + '.jpg'" type="image/jpeg" />
              <img
                class="card-img-top"
                :src="service.imagePath + '.jpg'"
                :alt="service.name"
              />
            </picture>
          </b-card-text>
          <template #footer>
            <a
              class="btn btn-sm btn-secondary"
              :href="'service/' + service.tag"
            >
              Подробнее
            </a>
          </template>
        </b-card>
      </b-card-group>
    </div>
    <div class="row">
      <!-- eslint-disable-next-line vue/no-v-html -->
      <div v-html="contents.bottomText"></div>
    </div>
  </div>
</template>

<script>
import axios from '~/plugins/axios'

export default {
  async asyncData({ route, error }) {
    const page = await axios.get('/page/' + encodeURIComponent(route.path))
    const rawContent = await axios.get('/content/' + page.data.id)
    const services = await axios.get('/service')
    const contents = rawContent.data.reduce(
      (a, x) => ({ ...a, [x.tag]: x.data }),
      {}
    )
    return {
      title: page.data.title,
      description: page.data.description,
      contents,
      services: services.data.filter((s) => s.parentId === 0),
    }
  },
  data() {
    return {
      title: '',
      description: '',
      contents: {},
      services: {},
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
