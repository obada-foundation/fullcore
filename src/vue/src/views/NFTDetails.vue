<template>
  <div>
    <div class="container">
      <div class="sp-box sp-shadow">
        <h3>NFT</h3>

        <div class="sp-type-form__field sp-form-group">
          DID: {{ nft.id }}
        </div>

        <div class="sp-type-form__field sp-form-group">
          URI: {{ nft.uri }}
        </div>

        <div class="sp-type-form__field sp-form-group">
          URI Hash: {{ nft.uri_hash }}
        </div>

        <h3>NFT Data</h3>

        <div class="sp-type-form__field sp-form-group">
          USN: {{ nft.data.usn }}
        </div>

        <div class="sp-type-form__field sp-form-group">
          Trust Anchor Token: {{ nft.data.trust_anchor_token }}
        </div>

        <div class="sp-type-form__field sp-form-group">
          Checksum: {{ nft.data.checksum }}
        </div>

        <h3>NFT Documents</h3>

        <table class="sp-blockstable">
          <thead>
            <tr>
              <th class="sp-blockstable__height"><strong>Document</strong></th>
              <th class="sp-blockstable__hash"><strong>Hash</strong></th>
            </tr>
          </thead>
          <tbody>
            <tr class="sp-blockdisplayline" v-if="nft.data.documents.length === 0">
              <td colspan="2">No Documents yet</td>
            </tr>
            <tr class="sp-blockdisplayline" v-for="doc in nft.data.documents" v-bind:key="doc.name">
              <td><a :href="doc.uri" target="_blank">{{ doc.name }}</a></td>
              <td>{{ doc.hash }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'NFTDetails',
  data() {
    return {
      nft: {
        data: {
          documents: []
        }
      }
    }
  },
  created() {
    this.initNFT()
  },
  methods: {
    initNFT() {
      let vm = this

      this.$store.dispatch("obadafoundation.fullcore.obit/QueryGetNft", {
        params: {
          did: this.$route.params.did,
        },
        options: {
          subscribe: true,
          all: true
        }
      }).then(function (resp) {
        vm.nft = resp
      });
    },
  },
}

</script>