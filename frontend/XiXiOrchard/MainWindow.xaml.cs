using System;
using System.Linq;
using System.Windows;
using System.Windows.Controls;
using OxyPlot;
using OxyPlot.Series;

namespace XiXiOrchard
{
    public partial class MainWindow : Window
    {
        public MainWindow()
        {
            InitializeComponent();
        }

        private void OnRunBacktestClick(object sender, RoutedEventArgs e)
        {
            var selectedStrategy = (ComboBoxItem)StrategySelector.SelectedItem;
            string strategyName = selectedStrategy.Content.ToString();

            var result = RunBacktest(strategyName);
            DisplayResult(result);
        }

        private double[] RunBacktest(string strategyName)
        {
            // 这里调用后端的策略模块进行回测，并返回价格数据
            // 目前为模拟数据
            return new double[] { 100, 102, 101, 105, 110, 108, 115 };
        }

        private void DisplayResult(double[] prices)
        {
            var model = new PlotModel { Title = "Backtest Result" };
            var lineSeries = new LineSeries { Title = "Price", MarkerType = MarkerType.Circle };

            for (int i = 0; i < prices.Length; i++)
            {
                lineSeries.Points.Add(new DataPoint(i, prices[i]));
            }

            model.Series.Add(lineSeries);
            ResultPlotView.Model = model;

            ResultTextBlock.Text = "Backtest completed. See the chart above for details.";
        }
    }
}
