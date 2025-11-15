import mongoose from 'mongoose';

const UserSchema = new mongoose.Schema({
  fullName: { type: String, required: true },
  email: { type: String, unique: true, required: true },
  role: { type: String, required: true },
  walletAddress: { type: String, required: true },
  faceVerified: { type: Boolean, default: false },
  mailVerified: { type: Boolean, default: false },
  otp: { type: String, default: null },
  selfieImage: { type: String, default: '' },
  otpExpiry: { type: Date, default: null },
  tags: { type: [String], default: [] },
  profileImage: { type: String, default: '' },
  rollNumber: { type: String, default: '' },
  program: { type: String, default: '' },
  major: { type: String, default: '' },
});

UserSchema.pre('save', function (next) {
  if (this.rollNumber) {
    const startYear = this.rollNumber.slice(0, 2);
    const batch = 2000 + parseInt(startYear) + 4; 
    const batchTag = `#Batch${batch}`;
    if (!this.tags.includes(batchTag)) this.tags.push(batchTag);
  }

  if (this.mailVerified && this.faceVerified) {
    if (!this.tags.includes('#verified')) this.tags.push('#verified');
  } else {
    this.tags = this.tags.filter((t: string) => t !== '#verified');
  }

  next();
});

delete mongoose.models.User;
export default mongoose.models.User || mongoose.model('User', UserSchema);
