import { Component } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { BillDetails, Profile, User,Name,Plan } from '../models/user';
import { UserService } from '../services/user.service';
import { decryptData } from '../helpers/aesEncryption';

@Component({
  selector: 'app-newprofile',
  templateUrl: './newprofile.component.html',
  styleUrls: ['./newprofile.component.scss'],
})
export class NewprofileComponent {
  encUid=localStorage.getItem('userID')|| '0'
  userId:string=decryptData(this.encUid)

  userData: User = new User('', new Name('',''), 0, '', '', [], new BillDetails('', new Plan('',0), '', ''));

  newProfile: Profile = new Profile('', '', '', [], [], [], []);
  profileForm: FormGroup;
  submitted: boolean = false;
  newID:string=''

  constructor(
    private router: Router,
    private formBuilder: FormBuilder,
    private userService: UserService,
    private snackbar: MatSnackBar
  ) {
    this.profileForm = this.formBuilder.group({
      profileName: ['', Validators.required],
    });
  }

  ngOnInit() {
    this.userService.getUserById(this.userId).subscribe((data) => {
      this.userData = data;
    });
  }
  get f() {
    return this.profileForm.controls;
  }
  onSubmit() {
    this.submitted = true;
    if (this.profileForm.invalid) {
      return;
    }

    if (this.userData.profile.length <= 4) {
      this.newID = this.generateId();
      var imageID = Math.floor(Math.random() * (5 - 1) + 1);
      var pName = this.profileForm.value.profileName;
      var flag = false;
      // console.log(pName);
      var avatar = '../../assets/img/avatars/avatar' + imageID + '.jpg';

      if (pName && avatar) {
        this.newProfile = new Profile(this.newID,pName,avatar,[],[],[],[]);
        this.userData.profile.push(this.newProfile);
        this.userService.updateUser(this.userData).subscribe((response) => {
          this.openSnackBar('New profile created successfully', 'Dismiss');
          this.router.navigate(['/profiles'])
        });
      }
    } else {
      this.openSnackBar("Maximum profiles created for this account","Dismiss")
    }
  }

  // To generate ID
    generateId() {
      const timestamp = Math.floor(new Date().getTime() / 1000).toString(16);
      return timestamp + 'xxxxxxxxxxxxxxxx'.replace(/[x]/g, () => {
        return (Math.floor(Math.random() * 16)).toString(16);
      }).toLowerCase();
    }
  
   openSnackBar(message: string, action: string) {
    this.snackbar.open(message, action, {
      duration:2000,
      panelClass: ['custom-snackbar', 'custom-snackbar'],
    });
  }
}
