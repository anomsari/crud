/**
 * @license
 * Copyright (c) 2016 The Polymer Project Authors. All rights reserved.
 * This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
 * The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
 * The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
 * Code distributed by Google as part of the polymer project is also
 * subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
 */

 import { PolymerElement, html } from '@polymer/polymer/polymer-element.js';
 import '@polymer/paper-input/paper-input.js';
 import '@polymer/iron-flex-layout/iron-flex-layout.js';
 import '@polymer/paper-item/paper-item.js';
 import '@polymer/iron-list/iron-list.js';
 import './shared-styles.js';

class MyNilai extends PolymerElement {
  static get template() {
    return html`
      <style include="shared-styles">
        :host {
          display: block;

          padding: 10px;
        }
        .vertical{
          @apply -- layout-vertical;
        }
        .horizontal{
          @apply -- layout-horizontal;
        }
        .kartu{
          margin: 24px;
          padding: 16px;
          color: white;
          border-radius: 5px;
          background-color: pink;
          box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 1px 5px 0 rgba(0, 0, 0, 0.12), 0 3px 1px -2px rgba(0, 0, 0, 0.2);

        }
        .kartu:hover{
          background-color: green;
          cursor : pointer;
        }
      </style>

      <div class="card">
        <div class="circle">1</div>
        <h1>Nilai</h1>

        <paper-input id="kd_mk" label="Kode Mata Kuliah" on-change="_change" readonly></paper-input>
        <paper-input id="npm" label="NPM" on-change="_change" readonly></paper-input>
        <paper-input id="uas" label="UAS" on-change="_change"readonly></paper-input>
        <paper-input id="uts" label="UTS" on-change="_change"readonly></paper-input>
        <paper-input id="total" label="Total" on-change="_change"readonly></paper-input>
        <paper-input id="ipk" label="IPK" on-change="_change"readonly></paper-input>

        <br>
        <!-- <label>NPM   : </label><label id="npmnya"></label><br>
        <label>uas  : </label><label id="uasnya"></label><br>
        <label>uts : </label><label id="utsnya"></label><br> -->

        <br>
        <!-- masukin data dummy pake list -->
        <iron-list items="[[mhs]]" as="item" on-selected-item-changed="_listChanged" selection-enabled>
          <template>
          <div>
          <div class="kartu" npm$="[[item.npm]]" >
            <div> no    :[[_index(index)]]</div>
            <div> kd_mk   : [[item.kd_mk]]</div>
            <div> npm   : [[item.npm]]</div>
            <div> uas  : [[item.uas]]</div>
            <div> uts : [[item.uts]]</div>
            <!-- <div> total : [[item.total]]</div>
            <div> ipk : [[item.ipk]]</div> -->
          </div>
          </div>
          </template>
        </iron-list>
        <!-- <p id="npmnya"></p>
        <p id="uasnya"></p>
        <p id="utsnya"></p> -->

      </div>
    `;
  }

  static get properties(){
    return{
      // nilai:{
      //   type    :Object,
      //   // notify  :true,
      //   value:{
      //     kd_mk:"",
      //     npm:"",
      //     // uas:"",
      //     // uts:"",
      //     // total:"",
      //     // ipk:"",
      //         }
      // },
      //contoh bikin data
      mhs:{
        type: Object,
        value:[
          {
            kd_mk:"1si2",
            npm   :"51418634",
            uas  :56,
            uts :86,

          },{
            kd_mk:"1ea2",
            npm   :"51418997",
            uas  :54,
            uts :75,

          },{
            kd_mk:"1pa2",
            npm   :"51418462",
            uas  :42,
            uts :54,

          },{
            kd_mk:"1st2",
            npm   :"51418451",
            uas  :85,
            uts :54,

          }
        ]
      }
    };


  }

  _change(e){
    let key = e.srcElement.id;
    this.nilai[e.srcElement.id]=e.target.value;
  console.log(this.nilai);
}

  _index(index){
    var a = parseInt(index+1);
     return a;
  }

  _listChanged(a){
    if (a.detail.value){

    var nilai = a.detail.value;
    let uts = nilai.uts;
    let uas = nilai.uas;
    let total = (nilai.uts+nilai.uas)/2;
    let ipk ="";

     if ( total> 80){
      ipk = "A";
    }else if ( total> 70){
      ipk = "B";
    }else if ( total> 60){
      ipk = "C";
    }else if(total > 50) {
      ipk = "D";
    }else{
      ipk = "P";
    }
    // nilai.value = nilai.ipk;



      this.$.kd_mk.value = nilai.kd_mk;
      this.$.npm.value = nilai.npm;
      this.$.uas.value = nilai.uas;
      this.$.uts.value = nilai.uts;
      this.$.total.value = total;
      this.$.ipk.value = ipk;

    }
  }
}

  //
  // _npm(e){
  //   // let npm = this.$.npmnya;
  //   // npm.innerHTML = e.target.value;
  //   this.$.npmnya.innerHTML = e.target.value;// manggil by id
  // }
  // _uas(e){
  //   // let npm = this.$.uasnya;
  //   // uas.innerHTML = e.target.value;
  //       this.$.uasnya.innerHTML = e.target.value;
  // }
  // _uts(e){
  //   // let npm = this.$.utsnya;
  //   // uts.innerHTML = e.target.value;
  //       this.$.utsnya.innerHTML = e.target.value;
  // }
  //


window.customElements.define('my-nilai', MyNilai);
