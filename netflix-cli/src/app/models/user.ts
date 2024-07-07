export class User {
  _id: string;
  name: Name;
  phone: number;
  email: string;
  password: string;
  profile: Profile[] = [];
  billingDetails: BillDetails;

  constructor(
    _id: string,
    name: Name,
    phone: number,
    email: string,
    password: string,
    profile: Profile[],
    billingDetails: BillDetails
  ) {
    this._id = _id;
    this.name = name;
    this.phone = phone;
    this.email = email;
    this.password = password;
    this.profile = profile;
    this.billingDetails = billingDetails;
  }
}

export class Name {
  first: string;
  last: string;

  constructor(first: string, last: string) {
    this.first = first;
    this.last = last;
  }
}

export class Profile {
  id: string;
  name: string;
  avatar: string;
  watchlist: number[] = [];
  watchhistory: number[] = [];
  liked: number[] = [];
  disliked: number[] = [];

  constructor(
    id: string,
    name: string,
    avatar: string,
    watchlist: number[],
    watchhistory: number[],
    liked: number[],
    disliked: number[]
  ) {
    this.id = id;
    this.name = name;
    this.avatar = avatar;
    this.watchlist = watchlist;
    this.watchhistory = watchhistory;
    this.liked = liked;
    this.disliked = disliked;
  }
}

export class BillDetails {
  id: string;
  plan: Plan;
  purchasedOn: string;
  cardNumber: string;

  constructor(
    id: string,
    plan: Plan,
    purchasedOn: string,
    cardNumber: string
  ) {
    this.id = id;
    this.plan = plan;
    this.purchasedOn = purchasedOn;
    this.cardNumber = cardNumber;
  }
}

export class Plan {
  name: string;
  price: number;

  constructor(name: string, price: number) {
    this.name = name;
    this.price = price;
  }
}
