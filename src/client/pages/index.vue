<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <div class="text-center">
        <alef-logo />
      </div>
      <v-card>
        <v-card-title class="headline">
          Lorem ipsum dolor sit amet.
        </v-card-title>
        <v-card-text>
          <p>
            Lorem, ipsum dolor sit amet consectetur adipisicing elit. Deleniti
            sapiente enim ab iure ut quidem non corrupti, quam possimus
            blanditiis pariatur itaque minus eum atque fugit repudiandae magnam.
            Accusamus, labore. Dignissimos rerum maxime ab expedita quam facere
            nesciunt ratione impedit temporibus! Perferendis cupiditate ex cum
            temporibus commodi recusandae consequatur voluptas, totam quasi
            praesentium, et dignissimos! Similique repellat libero commodi
            dolor?
          </p>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="primary" nuxt to="/inspire"> Continue </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import AlefLogo from '~/components/AlefLogo.vue'

export default {
  components: {
    AlefLogo,
  },
}
</script>
