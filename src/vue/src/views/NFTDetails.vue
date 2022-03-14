<template>
  <div>
    <div class="container">
      <div class="sp-component">
        <div class="sp-component-title">
          <h3>NFT</h3>
        </div>

        <div class="sp-box sp-shadow">
          <div class="sp-type-form__field sp-form-group text-wrap">
            <strong>DID:</strong><br />{{ nft.id }}
          </div>

          <div class="sp-type-form__field sp-form-group text-wrap">
            <strong>URI:</strong><br />{{ nft.uri }}
          </div>

          <div class="sp-type-form__field sp-form-group text-wrap">
            <strong>URI Hash:</strong><br />{{ shortHash(nft.uri_hash) }}
          </div>
        </div>
      </div>

      <div class="sp-component">
        <div class="sp-component-title">
          <h3>NFT Data</h3>
        </div>

        <div class="sp-box sp-shadow">
          <div class="sp-type-form__field sp-form-group text-wrap">
            <strong>USN:</strong><br />{{ nft.data.usn }}
          </div>

          <div class="sp-type-form__field sp-form-group text-wrap">
            <strong>Trust Anchor Token:</strong><br />{{ shortHash(nft.data.trust_anchor_token) }}
          </div>

          <div class="sp-type-form__field sp-form-group text-wrap">
            <strong>Checksum:</strong><br />{{ shortHash(nft.data.checksum) }}
          </div>
        </div>
      </div>

      <div class="sp-component">
        <div class="sp-component-title">
          <h3>NFT Documents</h3>
        </div>

        <div class="sp-box sp-shadow">
          <table class="sp-blockstable">
            <thead>
              <tr>
                <th class="sp-blockstable__hash"><strong>Document</strong></th>
                <th class="sp-blockstable__hash"><strong>Hash</strong></th>
              </tr>
            </thead>
            <tbody>
              <tr class="sp-blockdisplayline" v-if="nft.data.documents.length === 0">
                <td colspan="2">No Documents yet</td>
              </tr>
              <tr class="sp-blockdisplayline" v-for="doc in nft.data.documents" v-bind:key="doc.name">
                <td class="text-wrap">
                  <a :href="doc.uri" target="_blank">{{ doc.name }}</a>
                </td>
                <td class="text-wrap">
                  {{ shortHash(doc.hash) }}
                  <div class="sp-copy-icon" @click="copyHash(doc.hash)">
                    <span class="sp-icon sp-icon-Copy" />
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { copyToClipboard, shortenHash } from '@/utils/helpers'

export default {
  name: 'NFTDetails',
  data() {
    return {
      nft: {
        data: {
          documents: [],
        },
      },
    }
  },
  created() {
    this.initNFT()
  },
  methods: {
    initNFT() {
      let vm = this

      this.$store
        .dispatch('obadafoundation.fullcore.obit/QueryGetNft', {
          params: {
            did: this.$route.params.did,
          },
          options: {
            subscribe: true,
            all: true,
          },
        })
        .then(function (resp) {
          vm.nft = resp
        })
    },
    shortHash(string) {
      return shortenHash(string)
    },
    copyHash(string) {
      copyToClipboard(string)
    },
  },
}
</script>
