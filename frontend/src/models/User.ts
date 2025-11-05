// models/User.ts
import mongoose from 'mongoose';

const UserSchema = new mongoose.Schema({
  fullName: { type: String, required: true },
  email: { type: String, unique: true, required: true },
  role: { type: String, required: true },
  walletAddress: { type: String, required: true },
  faceVerified: { type: Boolean, default: false },
  mailVerified: { type: Boolean, default: false },
  otp: { type: String, default: null },
  otpExpiry: { type: Date, default: null },
  tags: { type: [String], default: [] }, 
});

UserSchema.pre('save', function (next) {
  if (this.mailVerified && this.faceVerified) {
    if (!this.tags.includes('#verified')) this.tags.push('#verified');
  } else {
    this.tags = this.tags.filter((t: string) => t !== '#verified');
  }
  next();
});

delete mongoose.models.User;
export default mongoose.models.User || mongoose.model('User', UserSchema);
