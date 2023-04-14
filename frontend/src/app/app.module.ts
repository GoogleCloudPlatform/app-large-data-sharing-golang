import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatChipsModule } from '@angular/material/chips';
import { AppComponent } from './app.component';
import { ListComponent } from './list/list.component';
import { UpdateFileComponent } from './update-file/update-file.component';
import { ViewComponent } from './view/view.component';
import { AppRoutingModule } from './app-routing.module';
import { allIcons, HeroIconModule } from 'ng-heroicon';
import { ConfirmDialogComponent } from './confirm-dialog/confirm-dialog.component';
import { FileUploadComponent } from './file-upload/file-upload.component';
import { OverlayModule } from '@angular/cdk/overlay';
import { HeaderComponent } from './header/header.component';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { SearchBarComponent } from './search-bar/search-bar.component';
import { SessionStorageService } from './service/session-storage.service';
import { ErrorInterceptor } from './interceptors/error-interceptor';
import { ErrorComponent } from './error/error.component';
@NgModule({
  declarations: [
    AppComponent,
    ListComponent,
    UpdateFileComponent,
    ViewComponent,
    ConfirmDialogComponent,
    FileUploadComponent,
    HeaderComponent,
    SearchBarComponent,
    ErrorComponent
  ],
  imports: [
    BrowserModule,
    OverlayModule,
    FormsModule,
    ReactiveFormsModule,
    MatSnackBarModule,
    BrowserAnimationsModule,
    MatChipsModule,
    HttpClientModule,
    AppRoutingModule,
    HeroIconModule.forRoot({
      ...allIcons
    }),
  ],
  providers: [
    SessionStorageService,
    { provide: HTTP_INTERCEPTORS, useClass: ErrorInterceptor, multi: true }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
