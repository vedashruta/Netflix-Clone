import { Component } from '@angular/core';
import { UserService } from '../services/user.service';
import { BillDetails, User,Name,Plan } from '../models/user';
import { Router } from '@angular/router';
import { decryptData, encryptData } from '../helpers/aesEncryption';

@Component({
  selector: 'app-profiles',
  templateUrl: './profiles.component.html',
  styleUrls: ['./profiles.component.scss']
})
export class ProfilesComponent {

  encId=localStorage.getItem('userID')|| '0'
  userId:string=decryptData(this.encId)

  userData: User = new User('', new Name('',''), 0, '', '', [], new BillDetails('', new Plan('',0), '', ''));

  constructor(private userService:UserService,private router:Router){

  }

  ngOnInit(){
    this.userService.getUserById(this.userId).subscribe((response)=>{
      this.userData=response
      }); 
  }

  selectProfile(profileId:string){
    var encprofileId:string=encryptData(profileId.toString())
    
    localStorage.setItem('profileID',encprofileId)
    window.location.replace('/home2')
  }

}


