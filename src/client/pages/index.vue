<template>
  <div class="container">
    <div class="row">
      <!-- eslint-disable-next-line vue/no-v-html -->
      <div v-html="contents.main"></div>
    </div>

    <div class="container mt-4">
      <div class="row justify-content-center">
        <div v-for="service of services" :key="service.id">
          <div class="col-auto mb-3">
            <div class="card" style="width: 18rem">
              <div class="card-body p-0">
                <h5 class="card-title p-2 m-0 text-center" style="height: 4rem">
                  {{ service.name }}
                </h5>
                <picture>
                  <source
                    :srcset="service.imagePath + '.webp'"
                    type="image/webp"
                  />
                  <source
                    :srcset="service.imagePath + '.jpg'"
                    type="image/jpeg"
                  />
                  <img
                    class="card-img-top"
                    :src="service.imagePath + '.jpg'"
                    :alt="service.name"
                  />
                </picture>
              </div>
              <div class="card-body p-3">
                <a
                  class="btn btn-sm btn-secondary"
                  :href="'service/' + service.tag"
                >
                  Подробнее
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
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
