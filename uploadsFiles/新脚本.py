# AUPR（Average Precision-Recall）
# FPR95（False Positive Rate at 95% Recall）
# ECE（Expected Calibration Error）

# AUPR和FPR95计算
aupr_scores = []
fpr95_scores = []
for i in range(len(CELL_TYPES)):
    binary_true = (all_labels == i).astype(np.int32)
    class_probs = all_probs[:, i]
    
    if np.any(binary_true):
        try:
            precision, recall, _ = precision_recall_curve(binary_true, class_probs)
            aupr_scores.append(auc(recall, precision))
            
            fpr, tpr, _ = roc_curve(binary_true, class_probs)
            if np.any(tpr >= 0.95):
                idx = np.argmin(np.abs(tpr - 0.95))
                fpr95_scores.append(fpr[idx])
        except Exception as e:
            print(f"类别 {i} 指标计算失败: {str(e)}")

metrics['AUPR'] = float(np.mean(aupr_scores) * 100) if aupr_scores else 0.0
metrics['FPR95'] = float(np.mean(fpr95_scores) * 100) if fpr95_scores else 0.0

# ECE计算
n_bins = 15
bin_boundaries = np.linspace(0, 1, n_bins + 1)
bin_confidences = np.max(all_probs, axis=1)
bin_accuracies = (all_preds == all_labels).astype(np.float32)
ece = 0.0
total_samples = len(bin_accuracies)

for bin_lower, bin_upper in zip(bin_boundaries[:-1], bin_boundaries[1:]):
    bin_mask = (bin_confidences >= bin_lower) & (bin_confidences < bin_upper)
    if np.any(bin_mask):
        bin_samples = np.sum(bin_mask)
        bin_acc = np.mean(bin_accuracies[bin_mask])
        bin_conf = np.mean(bin_confidences[bin_mask])
        ece += (bin_samples / total_samples) * np.abs(bin_acc - bin_conf)

metrics['ECE'] = float(ece * 100)

