import { NgModule } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatCardModule } from '@angular/material/card';
import {MatIconModule} from '@angular/material/icon';

const MaterialComponents = [
    MatButtonModule,
    MatInputModule,
    MatToolbarModule,
    MatFormFieldModule,
    MatCardModule,
    MatIconModule,
];

@NgModule({
    imports: [MaterialComponents],
    exports: [MaterialComponents]
})

export class MaterialModule { }