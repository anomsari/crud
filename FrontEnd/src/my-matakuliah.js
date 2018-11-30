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
 import '@polymer/paper-card/paper-card.js';
 import '@polymer/iron-flex-layout/iron-flex-layout.js';
 import '@polymer/paper-item/paper-item.js';
 import '@polymer/paper-item/paper-icon-item.js';
 import '@polymer/iron-list/iron-list.js';
 import '@polymer/iron-media-query/iron-media-query.js';
 import '@polymer/paper-fab/paper-fab.js';
 import './shared-styles.js';

class MyMatakuliah extends PolymerElement {
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
          border-collapse: collapse;
          width: 100%;
          text-align: left;
          padding: 2px;
        }

        .kartu:hover{
          background-color: green;
          cursor : pointer;
        }
        .kartu-mobile{
          border-collapse: collapse;
          width: 100%;
          text-align: left;
          padding: 2px;
        }
        .kartu-mobile:hover{
          background-color: red;
          cursor : pointer;
        }

        .paper-fab[label=Y]{
        front-weight :400px;
        front-size   :200px;
        }

        .paper-fab[label=Z]{
        front-weight :400px;
        front-size   :200px;
        }
        .margin{
          margin-left 10px;
        }
      </style>


      <div class="card">
        <div class="circle">1</div>
        <h1>Matakuliah</h1>
        <paper-item class="horizontal">
        <paper-input id="kd_mk" label="search" on-change="_change" readonly></paper-input>
        </paper-item>
      <br>
        <!-- <label>NPM   : </label><label id="npmnya"></label><br>
        <label>Nama  : </label><label id="namanya"></label><br>
        <label>Kelas : </label><label id="kelasnya"></label><br> -->
        <iron-media-query query="(min-width : 641px)" query-matches="{{desktop}}"> </iron-media-query>

        <br>
        <template is="dom-if" if="{{desktop}}">
        <!-- masukin data dummy pake list/ tampilan desktop -->
        <iron-list items="[[mhs]]" as="item" on-selected-item-changed="_listChanged" selection-enabled>
          <template>
          <div>
          <paper-item class="kartu horizontal" kd_mk$="[[item.kd_mk]]">

            <!-- <div style="width :20%"> [[_index(index)]]</div> -->
            <div style="width :50%"> [[item.kd_mk]]</div>
            <div style="width :50%"> [[item.mata_kuliah]]</div>
          </div>
          </div>
          </template>
        </iron-list>
        </template>


        <template is="dom-if" if="{{!desktop}}">
        <!-- masukin data dummy pake list/ tampilan hp -->
        <iron-list items="[[mhs]]" as="item" on-selected-item-changed="_listChanged" selection-enabled>
          <template>
          <div>
          <paper-item class="kartu-mobile horizontal" kd_mk$="[[item.kd_mk]]">
          <paper-fab label="[[_substring(item.kd_mk)]]">  </paper-fab>
          <div class="vertical margin">
          <!-- <div width :20%> [[_index(index)]]</div> -->
          <div style="width :50%"> [[item.kd_mk]]</div>
          <div style="width :50%"> [[item.mata_kuliah]]</div>
          </div>
          </div>
          </paper-item>
          </template>
        </iron-list>
        </template>
        <!-- <p id="npmnya"></p>
        <p id="namanya"></p>
        <p id="kelasnya"></p> -->
      </div>
     `;
   }

  static get properties(){
    return{
      matakuliah:{
        type    :Object,
        // notify  :true,
        value:{
          kd_mk:"",
          mata_kuliah:"",
        }
      },
      //contoh bikin data
      mhs:{
        type: Object,
        value:[
          {
            kd_mk   :"1si2",
            mata_kuliah  :"Sistem Informasi",

          },{
            kd_mk   :"1ea2",
            mata_kuliah  :"Sistem Data",
          },{
            kd_mk   :"1pa2",
            mata_kuliah  :"Sistem Terdistribusi",

          },{
            kd_mk   :"1st2",
            mata_kuliah  :"Algortima Pemrograman",

          }
        ]
      }
    };


  }
  _substring(s){

    var str = s.substring(1, 3);
    var sub = str.toUpperCase();
    return sub
  }
  _change(e){
    let key = e.srcElement.id;
    this.matakuliah[e.srcElement.id]=e.target.value;
  console.log(this.matakuliah);
}

  _index(index){
    var a = parseInt(index+1);
     return a
  }
  _listChanged(e){

    if (e.detail.value){
      let matakuliah = e.detail.value;
      this.$.kd_mk.value = matakuliah.kd_mk;
      this.$.mata_kuliah.value = matakuliah.mata_kuliah;
    }
  }
}

  //
  // _npm(e){
  //   // let npm = this.$.npmnya;
  //   // npm.innerHTML = e.target.value;
  //   this.$.npmnya.innerHTML = e.target.value;// manggil by id
  // }
  // _nama(e){
  //   // let npm = this.$.namanya;
  //   // nama.innerHTML = e.target.value;
  //       this.$.namanya.innerHTML = e.target.value;
  // }
  // _kelas(e){
  //   // let npm = this.$.kelasnya;
  //   // kelas.innerHTML = e.target.value;
  //       this.$.kelasnya.innerHTML = e.target.value;
  // }
  //


window.customElements.define('my-matakuliah', MyMatakuliah);
