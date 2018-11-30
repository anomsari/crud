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

class MyMahasiswa extends PolymerElement {
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

      <!-- <app-header-layout has-scro-religion style="width:100%; height:100vh;">
           <app-header class ="header" slot="header"fixed effects="waterfall">
           </app-header>         -->


      <div class="card">
        <div class="circle">1</div>
        <h1>Mahasiswa</h1>
        <paper-item class="horizontal">
        <paper-input id="npm" label="search" on-change="_change" ></paper-input>

        </paper-item>
        <br>
        <!-- <label>NPM   : </label><label id="npmnya"></label><br>
        <label>Nama  : </label><label id="namanya"></label><br>
        <label>Kelas : </label><label id="kelasnya"></label><br> -->

        <iron-media-query query="(min-width : 641px)" query-matches="{{desktop}}"> </iron-media-query>

        <br>
        <template is="dom-if" if="{{desktop}}">
        <!-- masukin data dummy pake list/ tampilan desktop-->
        <iron-list items="[[mhs]]" as="item" on-selected-item-changed="_listChanged" selection-enabled>
          <template>
          <div>
          <paper-item class="kartu horizontal" npm$="[[item.npm]]">

            <div style="width :20%"> [[item.npm]]</div>
            <div style="width :30%"> [[item.nama]]</div>
            <div style="width :30%"> [[item.kelas]]</div>
          </div>
          </paper-item>
          <!-- </div> -->
          </template>
        </iron-list>
        </template>


        <template is="dom-if" if="{{!desktop}}">
        <!-- masukin data dummy pake list/ tampilan hp -->
        <iron-list items="[[mhs]]" as="item" on-selected-item-changed="_listChanged" selection-enabled>
          <template>
          <div>
            <paper-item class="kartu-mobile horizontal" npm$="[[item.npm]]">
            <paper-fab label="[[_substring(item.nama)]]">  </paper-fab>
            <div class="vertical margin">
            <div style="width :20%"> [[item.npm]]</div>
            <div style="width :30%"> [[item.nama]]</div>
            <div style="width :30%"> [[item.kelas]]</div>
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
      mahasiswa:{
        type    :Object,
        // notify  :true,
        value:{
          npm:"",
          nama:"",
          kelas:"",
              }
      },
      //contoh bikin data
      mhs:{
        type: Object,
        value:[
          {
            npm   :"51418634",
            nama  :"Mona",
            kelas :"1IA01",
          },{
            npm   :"51418997",
            nama  :"Lisa",
            kelas :"1IA09",
          },{
            npm   :"51418462",
            nama  :"Agni",
            kelas :"1IA05",
          },{
            npm   :"51418451",
            nama  :"Asep",
            kelas :"1IA06",
          }
        ]
      }
    };


  }

  _substring(s){

    var str = s.substring(0, 1);
    var sub = str.toUpperCase();
    return sub
  }

  _change(e){
    let key = e.srcElement.id;
    this.mahasiswa[e.srcElement.id]=e.target.value;
  console.log(this.mahasiswa);
}

  _index(index){
    var a = parseInt(index+1);
     return a
  }
  _listChanged(e){

    if (e.detail.value){
      let mahasiswa = e.detail.value;
      this.$.npm.value = mahasiswa.npm;
      this.$.nama.value = mahasiswa.nama;
      this.$.kelas.value = mahasiswa.kelas;
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


window.customElements.define('my-mahasiswa', MyMahasiswa);
