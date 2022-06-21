<template>
    <b-container class="flex-column d-flex align-items-center">
        <b-form @submit="onSubmit">
            <b-col>
                <b-row class="">
                    <b-form-input class="mt-5" v-model="form.beans_amount" placeholder="Enter amout of beans"></b-form-input>
                </b-row>
            </b-col>

            <b-col>
                <b-row class="">
                    <b-form-input class="mt-5" v-model="form.beans_consumption" placeholder="Enter daily consumption"></b-form-input>
                </b-row>
            </b-col>

            <b-col>
                <b-row class="">
                    <b-form-select class="mt-5" v-model="form.shop_list" :options="shop_list" required></b-form-select>       
                </b-row>
            </b-col>            

            <b-col>
                <b-row class="">
                    <b-button class="mt-5" type="submit" variant="primary">{{btn_label}}</b-button>
                </b-row>
            </b-col>
         </b-form>
    </b-container>

    
</template>

<script>
  export default {
      // Registerしたら、登録した店舗のページへ遷移する
      // その時、RegisterボタンはEditに変更する
      // 削除ボタンもつけておく
    
    data() {
      return {        
        btn_label : 'Register',
        form : {
            beans_amount : this.beans_amount,
            beans_consumption :this.beans_consumption 
        },
        // TEST Data : labelはtext
        shop_list: [
            { text: 'Select One', value: null },
            { text:'shop1',value:1},
            { text:'shop2',value:2}
        ]
      }
    },
    methods : {
        onSubmit(event) {
            event.preventDefault()            
            // 下記で登録する
            //alert(JSON.stringify(this.form))            
            this.$router.push({
                        name:'editBeansInfo',
                        params: {
                            id: this.form.shop_list,
                            beans_amount : this.form.beans_amount,
                            beans_consumption : this.form.beans_consumption
                        }
            });
            
        }
    }

  }
</script>