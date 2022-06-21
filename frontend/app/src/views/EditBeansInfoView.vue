<template>

    <b-container class="flex-column d-flex align-items-center">
        <b-form @submit="onSubmit">
            <b-col>
                <b-row class="">
                    <b-form-input class="mt-5" v-model="form.beans_amount" placeholder="Enter amout of beans" v-bind:disabled="this.isDisabled">
                        {{this.$route.params.beans_amount}}
                        <!-- {{form.beans_amount}} -->
                    </b-form-input>
                </b-row>
            </b-col>

            <b-col>
                <b-row class="">
                    <b-form-input class="mt-5" v-model="form.beans_consumption" placeholder="Enter daily consumption" v-bind:disabled="this.isDisabled">
                        {{this.$route.params.beans_consumption}}
                        <!-- {{this.$route.params.beans_consumption}} -->
                    </b-form-input>
                </b-row>
            </b-col>

            <b-col>
                <b-row class="">                    
                    <b-form-select class="mt-5" v-model="form.shop" :options="shop_list" v-bind:disabled="this.isDisabled"></b-form-select>
                </b-row>
            </b-col>

            <b-col>
                <b-row class="">
                    <b-button class="mt-5" type="submit" variant="primary">{{this.btn_label}}</b-button>
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
        // console.log(beans_amount)
      return {
        isDisabled : true,
        btn_label : 'Edit', 
        form : {
            // DBで取得したデータをセットする
            beans_amount : this.$route.params.beans_amount,            
            beans_consumption : this.$route.params.beans_consumption,
            shop : this.$route.params.id
        },
        // DBからデータを取得する
        shop_list: [            
            { text:'shop1',value:1},
            { text:'shop2',value:2}
        ]   
      }
    },
    methods : {
        onSubmit(event) {
            event.preventDefault()
            if (this.isDisabled) {
                this.isDisabled = false
                this.btn_label = 'Update'
                return
            } else {
                this.isDisabled = true
                this.btn_label = 'Edit'
            }        
            
            // 更新処理
            alert('Update') 
            this.$router.push({
                        name:'editBeansInfo',
                        // params: {
                        //     id: this.form.shop,
                        //     beans_amount : this.form.beans_amount,
                        //     beans_consumption : this.form.beans_consumption
                        // }
            });
            // 下記で登録する
            //alert(JSON.stringify(this.form))
        }
    }

  }    
</script>