export enum Gender {
    Male = "Male",
    Female = "Female",
}

// UserProps holds User domain data
interface UserProps {
    ID?: string;
    name: string;
    age: number;
    gender: string; 
    email:string; 
    phone?: string;
}


export class User {
    private _ID?: string;
    private _name: string;
    private _age: number;
    private _gender: Gender
    private _email: string;
    private _phone?: string;

    private constructor(userData: UserProps) {
        this._name = userData.name;
        this._age = userData.age;
        this._gender = userData.gender as Gender;
        this._email = userData.email;
        this._phone = userData.phone;
        this._ID = userData.ID;
    }

    private static isValid(userData: UserProps): string[] {
        const errors: string[]= []

        if (!userData.name || userData.name.length < 2 || userData.name.length > 20){
            errors.push("Invalid Name");
        }

        if (!userData.age || userData.age <= 0 || userData.age > 150){
            errors.push("Invalid age");
        }

        if (userData.gender !== Gender.Male && userData.gender !== Gender.Female) {
            errors.push("Invalid gender");
        }

        if (!userData.email) { //TODO-> email validation
            errors.push("Invalid email");
        }
        
        console.log(errors);
        return errors;
    }

    public static NewUser(userData: UserProps): User {
        const errors = this.isValid(userData);
        if (errors.length !== 0) {
            throw new Error("Invalid data")
        }

        return new User(userData);
    }

    set ID(ID: string | undefined) {
        this._ID = ID;
    }

    get ID(): string | undefined {
        return this._ID;
    }

    get name(): string {
        return this._name;
    }

    get age(): number {
        return this._age;
    }

    get gender(): Gender {
        return this._gender;
    }

    get email(): string {
        return this._email;
    }

    get phone(): string | undefined {
        return this._phone;
    }
}
